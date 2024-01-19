package tls

import (
	"math/big"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/interfaces"
)

type CACert struct {
	Serial        *big.Int    `yaml:"serial"`
	ValidForYears int         `yaml:"validForYears"`
	Subject       CertSubject `yaml:"subject"`
}
type Cert struct {
	Serial        *big.Int    `yaml:"serial"`
	ValidForYears int         `yaml:"validForYears"`
	Subject       CertSubject `yaml:"subject"`
	DNSNames      []string    `yaml:"dnsNames"`
}
type CertSubject struct {
	Country            string `yaml:"country"`
	Organization       string `yaml:"organization"`
	OrganizationalUnit string `yaml:"organizationalUnit"`
	Locality           string `yaml:"locality"`
	Province           string `yaml:"province"`
	StreetAddress      string `yaml:"streetAddress"`
	PostalCode         string `yaml:"postalCode"`
	SerialNumber       string `yaml:"serialNumber"`
	CommonName         string `yaml:"commonName"`
}

var Irsa interfaces.Irsa = &interfaces.Rsa{}
var Ipem interfaces.Ipem = &interfaces.Pem{}
var Ix509 interfaces.Ix509 = &interfaces.X509{}
