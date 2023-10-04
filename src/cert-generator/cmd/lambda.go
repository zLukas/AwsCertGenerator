package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/aws"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

type RequestEvent struct {
	CACert tls.CACert `json:"caCert"`
	Cert   tls.Cert   `json:"cert"`
}

func handleRequest(ctx context.Context, event RequestEvent) (string, error) {

	caKey, ca, err := tls.CreateCACertBytes(&event.CACert)
	if err != nil {
		return "fail", fmt.Errorf("Failed to create CaCert: %s", err.Error())
	}
	ceKey, ce, err := tls.CreateCertBytes(&event.Cert, caKey, ca)
	if err != nil {
		return "fail", fmt.Errorf("Failed to create Cert: %s", err.Error())
	}
	//dbTable := os.Getenv("TABLE_NAME")
	dbTable := "CertTable"
	db := aws.Database{}
	err = db.PutItem(aws.TableRecord{
		CaCert: aws.CertItem{PrivateKey: string(caKey),
			Cert: string(ca),
		},
		CeCert: aws.CertItem{PrivateKey: string(ceKey),
			Cert: string(ce),
		},
		Name:         "sample-record",
		CreationDate: "today",
	},
		aws.WithDynamoDBLogin(),
		aws.WithTableName(dbTable),
	)

	return "sucess", nil

}

func RunLambda() {
	lambda.Start(handleRequest)
}
