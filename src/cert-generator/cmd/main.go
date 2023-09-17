package main

import (
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

// based on:  github.com/wardviaene/golang-for-devops-course/

//	type CertSubject struct {
//		Country            string `yaml:"country"`
//		Organization       string `yaml:"organization"`
//		OrganizationalUnit string `yaml:"organizationalUnit"`
//		Locality           string `yaml:"locality"`
//		Province           string `yaml:"province"`
//		StreetAddress      string `yaml:"streetAddress"`
//		PostalCode         string `yaml:"postalCode"`
//		SerialNumber       string `yaml:"serialNumber"`
//		CommonName         string `yaml:"commonName"`
//	}

// caCert:
//   serial: 1
//   validForYears: 10
//   subject:
//     country: US
//     organization: Golang Demo Org
//     organizationalUnit: Certificate Management
//     locality: NY
//     commonName: CA Certificate
// certs:
//   go-demo.localtest.me:
//     serial: 1
//     validForYears: 1
//     dnsNames: ["go-demo.localtest.me", "go-demo-2.localtest.me"]
//     subject:
//       country: US
//       organization: Golang Demo Org
//       organizationalUnit: go-demo department
//       locality: NY
//       commonName: go-demo.localtest.me
//   go-demo-client.localtest.me:
//     serial: 1
//     validForYears: 1
//     subject:
//       country: US
//       organization: Golang Demo Org
//       organizationalUnit: go-demo department
//       locality: NY
//       commonName: go-demo-client.localtest.me

func main() {

	ca := &tls.CACert{
		Serial:        1,
		ValidForYears: 1,
		Subject: tls.CertSubject{
			Country:            "PL",
			Organization:       "chmurPol",
			OrganizationalUnit: "dzial certow",
			Locality:           "PL",
			CommonName:         "certhost.jp2",
		},
	}

	ce := &tls.Cert{
		Serial:        1,
		ValidForYears: 1,
		dnsNames: ["yellowhost.jp2"],
		Subject: tls.CertSubject{
			Country:            "PL",
			Organization:       "chmurPol",
			OrganizationalUnit: "dzial rzultej chmury",
			Locality:           "PL",
			CommonName:         "yellowhost.jp2",
		},
	}
	caKey, caCert, err = tls.CreateCACert(&ca)
	Key, Cert, err = tls.CreateCert(&ce)



}
