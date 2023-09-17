package main

import (
	"fmt"
	"math/big"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func main() {
	ca := tls.CACert{
		Serial:        big.NewInt(1),
		ValidForYears: 1,
		Subject: tls.CertSubject{
			Country:            "PL",
			Organization:       "chmurPol",
			OrganizationalUnit: "dzial certow",
			Locality:           "PL",
			CommonName:         "certhost.jp2",
		},
	}

	ce := tls.Cert{
		Serial:        big.NewInt(1),
		ValidForYears: 1,
		DNSNames:      []string{"yellowhost.jp2"},
		Subject: tls.CertSubject{
			Country:            "PL",
			Organization:       "chmurPol",
			OrganizationalUnit: "dzial rzultej chmury",
			Locality:           "GD",
			CommonName:         "yellowhost.jp2",
		},
	}
	caKey, caCert, err := tls.CreateCACert(&ca)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	tls.WriteKeyCertFile(caKey, caCert, "CACert.pem")

	Key, Cert, err := tls.CreateCert(&ce, caKey, caCert)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	tls.WriteKeyCertFile(Key, Cert, "Cert.pem")
}
