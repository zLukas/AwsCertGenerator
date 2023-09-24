package tests

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

type mockPemOK struct{}
type mockPemFail struct{}
type mockX509OK struct{}
type mockX509Fail struct{}

func (m *mockPemOK) Decode(input []byte) (*tls.Block, []byte) {
	b := tls.Block{Bytes: input}
	return &b, nil
}
func (m *mockPemOK) Encode(out io.Writer, b *pem.Block) error {
	return nil
}

func (m *mockPemFail) Decode(input []byte) (*tls.Block, []byte) {
	return nil, nil
}

func (m *mockPemFail) Encode(out io.Writer, b *pem.Block) error {
	return fmt.Errorf("cannot encode buffer")
}

func (m *mockX509OK) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	b := []byte{0x56, 0xAA, 0x21}
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
