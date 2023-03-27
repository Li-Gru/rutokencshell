package main

import (
	"os"
)

type Config struct {
	selector   string
	pkcs11Id   string
	pkcs11Lib  string
	pkcs11Pin  string
	snxGateway string
	snxPrefix  string
	snxRealm   string
	snxPath    string
}

func defaultString(key string, val string) string {
	if key != "" || len(key) > 0 {
		return key
	}
	return val
}
func (c *Config) fromArg() {
	// TODO
}

func (c *Config) fromEnv() {
	c.selector = os.Getenv("CSHELL_PKCS11_SELECTOR")
	c.pkcs11Id = os.Getenv("CSHELL_PKCS11_ID")
	c.pkcs11Lib = os.Getenv("CSHELL_PKCS11_LIB")
	c.pkcs11Pin = os.Getenv("CSHELL_PKCS11_PIN")
	c.snxGateway = os.Getenv("CSHELL_SNX_GATEWAY")
	c.snxPrefix = os.Getenv("CSHELL_SNX_PREFIX")
	c.snxRealm = os.Getenv("CSHELL_SNX_REALM")
	c.snxPath = os.Getenv("CSHELL_SNX_PATH")

}

func (c *Config) defaults() {
	c.selector = defaultString(c.selector, "")
	c.snxPrefix = defaultString(c.snxPrefix, "/")
	c.snxPath = defaultString(c.snxPath, "/usr/bin/snx")
}

func (c *Config) Init() *Config {
	c.fromArg()
	c.fromEnv()
	c.defaults()
	return c
}
