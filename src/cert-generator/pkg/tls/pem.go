package tls

import (
	"crypto/x509"
	"fmt"
)

type Block struct{
	Bytes []byte
}

type IPem interface {
	Decode(data []byte) (*Block, []byte)
}


type Ix509 interface {
	
}

func PemToX509(input []byte, p IPem) (*x509.Certificate, error) {
	block, _ := p.Decode(input)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}	
	return x509.ParseCertificate(block.Bytes)
}	
