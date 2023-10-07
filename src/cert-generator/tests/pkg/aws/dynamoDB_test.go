package tests

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/zLukas/CloudTools/src/cert-generator/pkg/aws"
)

var test_table_record = aws.TableRecord{
	CaCert: aws.CertItem{PrivateKey: "CAPRIVKEY",
		Cert: "CACERT",
	},
	CeCert: aws.CertItem{PrivateKey: "CEPRIVKEY",
		Cert: "CECERT",
	},
	Name:         "sample-table",
	CreationDate: "today",
}

type WrongDbClient struct{}

func TestPutItemNilClient(t *testing.T) {
	test_db := aws.Database{TableName: "test-name"}
	err := test_db.PutItem(test_table_record)
	if err == nil {
		t.Error("err should be error type, got nil")
	} else if err.Error() != "database client is nil" {
		t.Errorf("err should be \"database client is nil\", got %s", err.Error())
	}
}

func TestPutItemNodynamoDBClient(t *testing.T) {
	test_db := aws.Database{Client: &WrongDbClient{}}
	err := test_db.PutItem(test_table_record)
	if err == nil {
		t.Error("err should be error type, got nil")
	} else if err.Error() != "client is not '*DynamoDB' type" {
		t.Errorf("err should be \"client is not '*DynamoDB' type\", got %s", err.Error())
	}

}

func TestPutItemNoTableName(t *testing.T) {
	test_db := aws.Database{Client: &dynamodb.DynamoDB{}}
	err := test_db.PutItem(test_table_record)
	if err == nil {
		t.Error("err should be error type, got nil")
	} else if err.Error() != "no table name provided" {
		t.Errorf("err should be \"no table name provided\", got %s", err.Error())
	}
}
