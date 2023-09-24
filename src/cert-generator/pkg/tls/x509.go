package tls

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
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

func CreateCACert(ca *CACert, p IPem, x Ix509) ([]byte, []byte, error) {
	template := &x509.Certificate{
		SerialNumber: ca.Serial,
		Subject: pkix.Name{
			Country:            removeEmptyString([]string{ca.Subject.Country}),
			Organization:       removeEmptyString([]string{ca.Subject.Organization}),
			OrganizationalUnit: removeEmptyString([]string{ca.Subject.OrganizationalUnit}),
			Locality:           removeEmptyString([]string{ca.Subject.Locality}),
			Province:           removeEmptyString([]string{ca.Subject.Province}),
			StreetAddress:      removeEmptyString([]string{ca.Subject.StreetAddress}),
			PostalCode:         removeEmptyString([]string{ca.Subject.PostalCode}),
			CommonName:         ca.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(ca.ValidForYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	keyBytes, certBytes, err := createCert(template, nil, nil, p, x)
	if err != nil {
		return nil, nil, err
	}

	return keyBytes, certBytes, nil
}

func CreateCert(cert *Cert, caKey []byte, caCert []byte, p IPem, x Ix509) ([]byte, []byte, error) {
	template := &x509.Certificate{
		SerialNumber: cert.Serial,
		Subject: pkix.Name{
			Country:            removeEmptyString([]string{cert.Subject.Country}),
			Organization:       removeEmptyString([]string{cert.Subject.Organization}),
			OrganizationalUnit: removeEmptyString([]string{cert.Subject.OrganizationalUnit}),
			Locality:           removeEmptyString([]string{cert.Subject.Locality}),
			Province:           removeEmptyString([]string{cert.Subject.Province}),
			StreetAddress:      removeEmptyString([]string{cert.Subject.StreetAddress}),
			PostalCode:         removeEmptyString([]string{cert.Subject.PostalCode}),
			CommonName:         cert.Subject.CommonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(cert.ValidForYears, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature,
		DNSNames:    removeEmptyString(cert.DNSNames),
	}

	caKeyParsed, err := key.PrivateKeyPemToRSA(caKey)
	if err != nil {
		return nil, nil, err
	}
	caCertParsed, err := PemToX509(caCert, p, x)
	if err != nil {
		return nil, nil, err
	}

	keyBytes, certBytes, err := createCert(template, caKeyParsed, caCertParsed, p, x)
	if err != nil {
		return nil, nil, err
	}
	return keyBytes, certBytes, nil
}

func createCert(template *x509.Certificate, caKey *rsa.PrivateKey, caCert *x509.Certificate, p IPem, x Ix509) ([]byte, []byte, error) {
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
		derBytes, err = x.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
		if err != nil {
			return nil, nil, err
		}
	} else {
		derBytes, err = x.CreateCertificate(rand.Reader, template, caCert, &privateKey.PublicKey, caKey)
		if err != nil {
			return nil, nil, err
		}
	}

	if err = p.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, nil, err
	}
	if err = p.Encode(&keyOut, key.RSAPrivateKeyToPEM(privateKey)); err != nil {
		return nil, nil, err
	}

	return keyOut.Bytes(), certOut.Bytes(), nil
}

func removeEmptyString(input []string) []string {
	if len(input) == 1 && input[0] == "" {
		return []string{}
	}
	return input
}
