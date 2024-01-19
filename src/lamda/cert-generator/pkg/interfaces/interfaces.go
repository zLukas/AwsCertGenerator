package interfaces

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

type Irsa interface {
	GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error)
}

type Ix509 interface {
	CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error)
	ParseCertificate(der []byte) (*x509.Certificate, error)
}

type Ipem interface {
	Decode(data []byte) (*pem.Block, []byte)
	Encode(out io.Writer, b *pem.Block) error
}

type Rsa struct{}

func (r *Rsa) GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(random, bits)
}

type Pem struct{}

func (p *Pem) Decode(data []byte) (*pem.Block, []byte) {
	return pem.Decode(data)
}

func (p *Pem) Encode(out io.Writer, b *pem.Block) error {
	return pem.Encode(out, b)
}

type X509 struct{}

func (x *X509) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, pub, priv)
}

func (x *X509) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(der)
}

//x509.MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
