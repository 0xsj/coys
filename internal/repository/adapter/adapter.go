package adapter

import "github.com/aws/aws-sdk-go/service/dynamodb"

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

func (db *Database) FindOne() {}

func (db *Database) CreateOrUpdate() {}

func (db *Database) Delete() {}

func (db *Database) Health() bool {
	return false
}
