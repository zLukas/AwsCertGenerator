package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Database struct {
	TableName string
	Client    interface{}
}

type DatabaseOption func(*Database)

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
