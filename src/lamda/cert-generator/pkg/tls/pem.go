package tls

import (
	"crypto/x509"
	"fmt"
)

func PemToX509(input []byte) (*x509.Certificate, error) {
	block, _ := Ipem.Decode(input)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return Ix509.ParseCertificate(block.Bytes)
}
