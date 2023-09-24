package tls

import (
	"crypto/x509"
	"fmt"
)

func PemToX509(input []byte, p IPem, x Ix509) (*x509.Certificate, error) {
	block, _ := p.Decode(input)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return x.ParseCertificate(block.Bytes)
}
