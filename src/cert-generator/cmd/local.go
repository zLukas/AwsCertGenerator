package cmd

import (
	"fmt"
	"os"

	"github.com/zLukas/CloudTools/src/cert-generator/pkg/aws"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/input"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

func RunLocal() {
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

	for k, el := range config.Cfg.Cert {

		ceKey, ce, err := tls.CreateCertBytes(el, caKey, ca)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		tls.WriteKeyCertFile(ceKey, ce, k+".pem")
	}

	tls.WriteKeyCertFile(caKey, ca, "CA-Certificate.pem")

	fmt.Print("uploading to database...")
	dbTable := "Certificates"
	db := aws.Database{}
	err = db.PutItem(aws.TableRecord{
		CaCert: aws.CertItem{PrivateKey: string(caKey),
			Cert: string(ca),
		},
		CeCert:       aws.CertItem{},
		Name:         "sample-record",
		CreationDate: "today",
	},
		aws.WithDynamoDBLogin(),
		aws.WithTableName(dbTable),
	)
	if err != nil {
		fmt.Printf("database upload error: %s", err.Error())
	}
}
