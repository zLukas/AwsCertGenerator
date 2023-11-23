//go:build awslambda

package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/aws"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/tls"
)

type RequestEvent struct {
	CACert    tls.CACert `json:"caCert"`
	Cert      tls.Cert   `json:"cert"`
	Requester string     `json:"requester"`
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
	dbTable := os.Getenv("TABLE_NAME")
	dbRegion := os.Getenv("DB_REGION")
	db := aws.Database{}
	if err != nil {
		return "fail", fmt.Errorf("Error: %s", err)
	}
	currentTime := time.Now()

	err = db.PutItem(aws.TableRecord{
		CaCert: aws.CertItem{
			PrivateKey: caKey,
			Cert:       ca,
		},
		CeCert: aws.CertItem{
			PrivateKey: ceKey,
			Cert:       ce,
		},
		Name:         event.Requester,
		CreationDate: currentTime.Format("2006.01.02 15:04:05"),
	},
		aws.WithDynamoDBLogin(dbRegion),
		aws.WithTableName(dbTable),
	)
	if err != nil {
		return "fail", fmt.Errorf("database upload error: %s", err.Error())
	}

	return "sucess", nil

}

func Run() {
	lambda.Start(handleRequest)
}
