package tests

import (
	"crypto/x509"
	"fmt"
	"io"
)

type mockCreateCertificateCA struct{}
type mockCreateCertificateCE struct{}

func (m *mockCreateCertificateCA) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	b := []byte{0xAA}
	return b, nil
}

func (m *mockX509Fail) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	b := []byte{0x56, 0xAA, 0x21}
	return b, nil
}

func (m *mockX509OK) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return &x509.Certificate{}, nil
}

func (m *mockX509Fail) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return nil, fmt.Errorf("x509: malformed certificate")
}
