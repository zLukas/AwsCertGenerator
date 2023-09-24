package tls

import (
	"crypto/x509"
	"encoding/pem"
	"io"
	"math/big"
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

type Block struct {
	Bytes []byte
}

type IPem interface {
	Decode(data []byte) (*Block, []byte)
	Encode(out io.Writer, b *pem.Block) error
}

type Ix509 interface {
	ParseCertificate(der []byte) (*x509.Certificate, error)
	CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error)
}
