package main

import (
	"fmt"
	"math/big"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func main() {
	generator := Generator{}
	caTemplate := tls.CACert{
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

	ceTemplate := tls.Cert{
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

	caKey, ca, err := tls.CreateCACertBytes(&caTemplate, &generator, &generator)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	ceKey, ce, err := tls.CreateCertBytes(&ceTemplate, caKey, ca, &generator, &generator)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	out := Output{}
	yaml_cert, err := out.format(Cert{
		Name:        "CA",
		Generated:   "today",
		Certificate: string(ca[:]),
		Key:         string(caKey[:]),
	}, Cert{
		Name:        "CE",
		Generated:   "today",
		Certificate: string(ce[:]),
		Key:         string(ceKey[:]),
	},
	)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(yaml_cert[:]))
}
