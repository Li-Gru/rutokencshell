package main

func main() {
	config := (&Config{}).Init()
	token := (&Token{}).init(&config.pkcs11Lib, &config.pkcs11Id, &config.pkcs11Pin)
	certificate := token.getCertificate(config.selector)
	if certificate == nil {
		panic("no certs")
	}
	extender := (&Extender{}).Init(*certificate)
	params := extender.GetParams(config.snxGateway, config.snxPrefix, config.snxRealm)

	snx := &SNX{
		SnxPath: config.snxPath,
		Params:  *params,
		Debug:   false,
	}
	snx.generateSNXInfo()
	printJson(snx)
	snx.callSNX()
}

/*
how to use:

env \
	CSHELL_PKCS11_LIB="/usr/lib/librtpkcs11ecp.so" \
	CSHELL_PKCS11_ID="123c45b678" \
	CSHELL_PKCS11_PIN="12345678" \
	CSHELL_PKCS11_SELECTOR="example@localhost.local" \
	CSHELL_SNX_PREFIX="/sslvpn" \
	CSHELL_SNX_GATEWAY="ssl.vpn.localhost.local" \
	CSHELL_SNX_REALM="ssl_vpn" \
	./cshell
*/
