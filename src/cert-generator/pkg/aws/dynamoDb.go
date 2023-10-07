package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func WithDynamoDBLogin(region string) DatabaseOption {
	return func(t *Database) {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		cfg.Region = region
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

func dynamoDBPutItem(client *dynamodb.Client, item TableRecord, table string) error {
	_, err := client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(table)})
	if err != nil {
		return fmt.Errorf("table error: %s", err.Error())
	}

	dbItem, err := attributevalue.MarshalMap(item)

	if err != nil {
		return fmt.Errorf("cannot marshal item into dynamoDBFormat")
	}
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: dbItem})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return nil
}
