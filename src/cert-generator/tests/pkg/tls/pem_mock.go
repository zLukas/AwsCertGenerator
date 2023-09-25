package tests

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

type mockPemOK struct{}
type mockPemFail struct{}
type mockX509OK struct{}
type mockX509Fail struct{}

func (m *mockPemOK) Decode(input []byte) (*pem.Block, []byte) {
	b := pem.Block{Bytes: input}
	return &b, nil
}
func (m *mockPemOK) Encode(out io.Writer, b *pem.Block) error {
	return nil
}

func (m *mockPemFail) Decode(input []byte) (*pem.Block, []byte) {
	return nil, nil
}

func (m *mockPemFail) Encode(out io.Writer, b *pem.Block) error {
	return fmt.Errorf("cannot encode buffer")