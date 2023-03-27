package main

import (
	"crypto/tls"

	"github.com/ThalesIgnite/crypto11"
)

type Token struct {
	context *crypto11.Context
}

func (t *Token) init(library *string, serial *string, pin *string) *Token {
	config := &crypto11.Config{
		Path:        *library,
		TokenSerial: *serial,
		Pin:         *pin,
	}
	context, err := crypto11.Configure(config)
	iferr(err)
	t.context = context
	return t
}

func (t *Token) getCertificate(selector string) *tls.Certificate {
	certificates, err := t.context.FindAllPairedCertificates()
	iferr(err)
	if len(certificates) == 0 {
		return nil
	}

	if selector == "" {
		return &certificates[0]
	}

	for _, certificate := range certificates {
		if contains(certificate.Leaf.EmailAddresses, selector) {
			return &certificate
		}
	}
	return nil
}
