package tls

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/key"
)

func WriteKeyCertFile(Key []byte, Cert []byte, filePath string) error {
	CertKey := append(Cert, Key...)

	if err := os.WriteFile(filePath, CertKey, 0600); err != nil {
		return err
	}
	return nil
}

func CreateCACertBytes(ca *CACert) ([]byte, []byte, error) {

	if ca == nil {
		return nil, nil, fmt.Errorf("template is nil")
	}

	if ca.Subject.CommonName == "" {
		return nil, nil, fmt.Errorf("template must contain CommonName Field")
	}

	template := &x509.Certificate{
		SerialNumber: ca.Serial,
		Subject: pkix.Name{
			Country:            RemoveEmptyString([]string{ca.Subject.Country}),
			Organization:       RemoveEmptyString([]string{ca.Subject.Organization}),
			OrganizationalUnit: RemoveEmptyString([]string{ca.Subject.OrganizationalUnit}),
			Locality:           RemoveEmptyString([]string{ca.Subject.Locality}),
			Province:           RemoveEmptyString([]string{ca.Subject.Province}),
			StreetAddress:      RemoveEmptyString([]string{ca.Subject.StreetAddress}),
			PostalCode:         RemoveEmptyString([]string{ca.Subject.PostalCode}),
			CommonName:         ca.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(ca.ValidForYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	keyBytes, certBytes, err := createCert(template, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	return keyBytes, certBytes, nil
}

func CreateCertBytes(cert *Cert, caKey []byte, caCert []byte) ([]byte, []byte, error) {
	template := &x509.Certificate{
		SerialNumber: cert.Serial,
		Subject: pkix.Name{
			Country:            RemoveEmptyString([]string{cert.Subject.Country}),
			Organization:       RemoveEmptyString([]string{cert.Subject.Organization}),
			OrganizationalUnit: RemoveEmptyString([]string{cert.Subject.OrganizationalUnit}),
			Locality:           RemoveEmptyString([]string{cert.Subject.Locality}),
			Province:           RemoveEmptyString([]string{cert.Subject.Province}),
			StreetAddress:      RemoveEmptyString([]string{cert.Subject.StreetAddress}),
			PostalCode:         RemoveEmptyString([]string{cert.Subject.PostalCode}),
			CommonName:         cert.Subject.CommonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(cert.ValidForYears, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
		DNSNames:    RemoveEmptyString(cert.DNSNames),
	}

	caKeyParsed, err := key.PrivateKeyPemToRSA(caKey)
	if err != nil {
		return nil, nil, err
	}
	caCertParsed, err := PemToX509(caCert)
	if err != nil {
		return nil, nil, err
	}

	keyBytes, certBytes, err := createCert(template, caKeyParsed, caCertParsed)
	if err != nil {
		return nil, nil, err
	}
	return keyBytes, certBytes, nil
}

func createCert(template *x509.Certificate, caKey *rsa.PrivateKey, caCert *x509.Certificate) ([]byte, []byte, error) {
	var (
		derBytes []byte
		certOut  bytes.Buffer
		keyOut   bytes.Buffer
	)

	privateKey, err := key.CreateRSAPrivateKey(4096)
	if err != nil {
		return nil, nil, err
	}
	if template.IsCA {
		derBytes, err = Ix509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
		if err != nil {
			return nil, nil, err
		}
	} else {
		derBytes, err = Ix509.CreateCertificate(rand.Reader, template, caCert, &privateKey.PublicKey, caKey)
		if err != nil {
			return nil, nil, err
		}
	}

	if err = Ipem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, nil, err
	}
	if err = Ipem.Encode(&keyOut, key.RSAPrivateKeyToPEM(privateKey)); err != nil {
		return nil, nil, err
	}

	return keyOut.Bytes(), certOut.Bytes(), nil
}

func RemoveEmptyString(input []string) []string {
	if len(input) == 1 && input[0] == "" {
		return []string{}
	}
	return input
}
