package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
}

func NewAdapter() Interface {
	return nil
}

func (db *Database) FindAll() {}

func (db *Database) FindOne(condition map[string]interface{}, tablename string) (response *dynamodb.GetItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tablename),
		Key:       conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate() {}

func (db *Database) Delete() {}

func (db *Database) Health() bool {
	return false
}
