package instance

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// return a pointer to dynamodb struct
func GetConnection() *dynamodb.DynamoDB {
	session := session.Must(session.NewSessionWithOptions(session.Options{ // create new session with shared config enabld
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(session)
}
