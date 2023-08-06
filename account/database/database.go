package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	context context.Context
	address string
	client  *mongo.Client
	db      *mongo.Database
}
