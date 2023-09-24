package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

type Generator struct{}
type Output struct {
	CaCert Cert `yaml:"ca"`
	CeCert Cert `yaml:"cert"`
}

type Cert struct {
	Name        string `yaml:"name"`
	Generated   string `yaml:"generation-time"`
	Certificate string `yaml:"certificate"`
	Key         string `yaml:"private-key"`
}

func (g *Generator) Decode(data []byte) (*pem.Block, []byte) {
	return pem.Decode(data)
}

func (g *Generator) Encode(out io.Writer, b *pem.Block) error {
	return pem.Encode(out, b)
}

func (g *Generator) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(der)
}

func (g *Generator) CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub any, priv any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, pub, priv)
}

func (o *Output) format(ca Cert, ce Cert) ([]byte, error) {
	o.CaCert = ca
	o.CeCert = ce
	yamlData, err := yaml.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("cannot convert to yaml: %s", err)
	}
	return yamlData, nil
}
