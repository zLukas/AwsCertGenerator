package input

import (
	"fmt"
	"os"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
	"gopkg.in/yaml.v2"
)

type CertSpecs struct {
	CACert *tls.CACert          `yaml:"caCert"`
	Cert   map[string]*tls.Cert `yaml:"certs"`
}

type Config struct {
	Cfg         CertSpecs
	CfgFilePath string
}

func (c *Config) ParseInput() error {

	if c.CfgFilePath == "" {
		c.CfgFilePath = "tls.yaml"
	}
	cfgFileBytes, err := os.ReadFile(c.CfgFilePath)
	if err != nil {
		return fmt.Errorf("Error while reading config file: %s\n", err)
	}
	err = yaml.Unmarshal(cfgFileBytes, &c.Cfg)
	if err != nil {
		return fmt.Errorf("Error while parsing config file: %s\n", err)
	}
	return nil
}
