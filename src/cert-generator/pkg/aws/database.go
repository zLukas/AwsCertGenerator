package aws

import (
	"fmt"
	"log"

	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

type CertItem struct {
	PrivateKey []byte `dynamodbav:"privateKey"`
	Cert       []byte `dynamodbav:"cert"`
}

type TableRecord struct {
	CaCert       CertItem `dynamodbav:"ca"`
	CeCert       CertItem `dynamodbav:"ce"`
	Name         string   `dynamodbav:"name"`
	CreationDate string   `dynamodbav:"creationDate"`
}
type Database struct {
	TableName string
	Client    interface{}
}

type DatabaseOption func(*Database)

func WithDynamoDBLogin() DatabaseOption {
	return func(t *Database) {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		cfg.Region = "eu-central-1"
		if err != nil {
			fmt.Printf("Cannot log into DB: %s", err.Error())
			t.Client = nil
			return
		}
		client := dynamodb.NewFromConfig(cfg)
		if client == nil {
			fmt.Printf("Cannot log into DB: %s", err.Error())
			t.Client = nil
			return
		}
		t.Client = client
	}
}

func WithTableName(n string) DatabaseOption {
	return func(t *Database) {
		t.TableName = n
	}
}

func (t *Database) PutItem(item TableRecord, opts ...DatabaseOption) error {
	for _, opt := range opts {
		opt(t)
	}

	if t.Client == nil {
		return fmt.Errorf("database client is nil")
	}

	if dynamoDbClient, ok := t.Client.(*dynamodb.Client); ok {
		if err := dynamoDBPutItem(dynamoDbClient, item, t.TableName); err != nil {
			return fmt.Errorf("failed to put dynamoDB item: %s", err.Error())
		}
	} else {
		return fmt.Errorf("client not supported")
	}
	return nil
}

func dynamoDBPutItem(client *dynamodb.Client, item TableRecord, table string) error {
	_, err := client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(table)})
	if err != nil {
		return fmt.Errorf("table error: %s", err.Error())
	}

	dbItem, err := attributevalue.MarshalMap(&item)
	if err != nil {
		panic(err)
	}
	fmt.Println("in ", dbItem["name"], dbItem["ca"])
	fmt.Printf("table %s\n", table)
	put_out, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: dbItem,
	})
	fmt.Printf("out %v\n", put_out)
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return nil
}
