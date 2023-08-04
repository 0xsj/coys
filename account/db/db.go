package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	context context.Context
	address string
	client  *mongo.Client
	db      *mongo.Database
}

func Connect() {}

func Disconnect() {}
