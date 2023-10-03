package aws

import (
	"fmt"
	originAws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type CertItem struct {
	PrivateKey string
	Cert       string
}

type TableRecord struct {
	CaCert       CertItem
	CeCert       CertItem
	Name        string
	CreationDate string
}
type Table struct {
	Name string
	client interface
}

type TableOption func(*Table)

func WithLogin() TableOption {
	return func(t *Table) {
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})
		if err != nil {
			fmt.Printf("Cannot log into DB: %s", err.Error())
			t.client = nil
		}

		t.client = dynamodb.New(sess)
	}
}

func WithName(n string) TableOption {
	return func(t *Table) {
		t.Name = n
	}
}


func (t *Table) PutItem(item map[string]TableRecord, opts ...TableOption) error {
		for _, opt := range opts {
			opt(t)
		}

		if t.client != nil {
			if _, ok = t.client.(*DynamoDB); ok != false {
				return fmt.Errorf("client is not '*DynamoDB' type")
			}
		} else {
			return fmt.Errorf("client is not '*DynamoDB' type")
		}

		var err error
		var tableItem map[string]*dynamodb.AttributeValue
		tableItem, err = dynamodbattribute.MarshalMap(item)
		if err != nil {
			return fmt.Errorf("failed to Marshal item %s", err.Error())
		}

		input := &dynamodb.PutItemInput{
			Item:      tableItem,
			TableName: originAws.String(t.Name),
		}

		_, err = t.client.PutItem(input)
		if err != nil {
			return fmt.Errorf("failed to put %d item %s", idx, err.Error())
		}
	return nil
}
