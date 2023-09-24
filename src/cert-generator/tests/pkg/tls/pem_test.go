package test

import (
	"testing"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

type mockPemOK struct {}
type mockPemFail struct {}

func (m *mockPemOK) Decode(input []byte) (*tls.Block, []byte){
	b := tls.Block{Bytes: []byte{0xAA, 0xC5, 0xAB}}
	return &b, nil
}

func (m *mockPemFail) Decode(input []byte) (*tls.Block, []byte){
	return nil, nil
}


func TestPemToX509_fail(t *testing.T) {
	var pemMock = mockPemOK{}
	var false_bytes = []byte{0xAA, 0xC5, 0xAB}
	results, err := tls.PemToX509(false_bytes, &pemMock)
	if results != nil {
		t.Errorf("results var execept to be nil, got %v ", results)
	}
	if err == nil {
		t.Errorf("results var execept to be nil, got %s ", err)
	}

}
