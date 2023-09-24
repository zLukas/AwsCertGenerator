package tests

import (
	"testing"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func TestPemToX509_ok(t *testing.T) {
	pemMock := mockPemOK{}
	x509Mock := mockX509OK{}
	var false_bytes = []byte{0xAA, 0xC5, 0xAB}
	results, err := tls.PemToX509(false_bytes, &pemMock, &x509Mock)
	if err != nil {
		t.Errorf("err expected to be nil, got %s ", err)
	}
	if results == nil {
		t.Errorf("results var execept to be %v, got nil ", tls.Block{Bytes: false_bytes})
	}

}

func TestPemToX509_fail(t *testing.T) {
	pemMock := mockPemFail{}
	x509Mock := mockX509Fail{}
	var false_bytes = []byte{0xAA, 0xC5, 0xAB}
	results, err := tls.PemToX509(false_bytes, &pemMock, &x509Mock)
	if results != nil {
		t.Errorf("results var execept to be nil, got %v ", results)
	}
	if err == nil {
		t.Errorf("results var execept to be nil, got %s ", err)
	}

}
