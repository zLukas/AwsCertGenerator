package tls

import (
	"testing"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func TestPemToX509_fail(t *testing.T) {
	var false_bytes = []byte{0xAA, 0xC5, 0xAB}
	results, err := tls.PemToX509(false_bytes)
	if results != nil {
		t.Errorf("results var execept to be nil, got %v ", results)
	}
	if err == nil {
		t.Errorf("results var execept to be nil, got %s ", err)
	}

}
