package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
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
	fmt.Printf("%v,%v", ceKey, ce)
	return "sucess", nil

}

func RunLambda() {
	lambda.Start(handleRequest)
}
