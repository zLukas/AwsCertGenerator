//go:build !awslambda

package cmd

import (
	"fmt"
	"os"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/input"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func Run() {
	var file string = ""
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	config := input.Config{CfgFilePath: file}

	if err := config.ParseInput(); err != nil {
		fmt.Printf("cannot parse input file: %s", err.Error())
		return
	}

	caKey, ca, err := tls.CreateCACertBytes(config.Cfg.CACert)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	tls.WriteKeyCertFile(caKey, ca, "CA-Certificate.pem")

}
