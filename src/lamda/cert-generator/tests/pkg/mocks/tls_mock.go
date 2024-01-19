package mocks

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

var testError error = fmt.Errorf("FAIL")

type X509Mock struct {
	createcertificatesPass bool
	parseCertificatePass   bool
}

func (x *X509Mock) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	if x.createcertificatesPass {
		return []byte{0xDE, 0xAD, 0xBE, 0xEF}, nil
	} else {
		return nil, testError
	}
}

func (x *X509Mock) ParseCertificate(der []byte) (*x509.Certificate, error) {
	if x.parseCertificatePass {
		return &x509.Certificate{}, nil
	} else {
		return nil, testError
	}
}

type PemMock struct {
	encodePass bool
	decodePass bool
}

func (i *pemMock) Encode(out io.Writer, b *pem.Block) error {
	if i.encodePass {
		return nil
	} else {
		return testError
	}
}

func (i *pemMock) Decode(data []byte) (*pem.Block, []byte) {
	fmt.Print(("Entering decode mock"))
	if i.decodePass == true {
		fmt.Print("returning pass")
		return &pem.Block{Type: "test", Bytes: []byte{0xDE, 0xAD, 0xBE, 0xEF}}, nil
	} else {
		fmt.Print("returning fail")
		return nil, nil
	}
}

type rsaMock struct {
	generateKeyPass bool
}

func (r *rsaMock) GenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error) {
	if r.generateKeyPass {
		return &rsa.PrivateKey{}, nil
	} else {
		return nil, testError
	}
}
