package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/interfaces"
)

var Irsa interfaces.Irsa = &interfaces.Rsa{}
var Ipem interfaces.Ipem = &interfaces.Pem{}
var Ix509 interfaces.Ix509 = &interfaces.X509{}

func CreateRSAPrivateKey(n int) (*rsa.PrivateKey, error) {
	return Irsa.GenerateKey(rand.Reader, n)
}

func RSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) *pem.Block {
	return &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
}

func CreateRSAPrivateKeyAndSave(path string, n int) error {
	privateKey, err := CreateRSAPrivateKey(n)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	if err := Ipem.Encode(f, RSAPrivateKeyToPEM(privateKey)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func PrivateKeyPemToRSA(input []byte) (*rsa.PrivateKey, error) {
	var parsedKey *rsa.PrivateKey
	var err error

	privPem, _ := Ipem.Decode(input)

	if privPem.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("RSA private key is of the wrong type: %s", privPem.Type)
	}

	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPem.Bytes); err != nil {
		return nil, fmt.Errorf("Unable to parse RSA private key: %v", err)
	}

	return parsedKey, nil
}
