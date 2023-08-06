package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	context context.Context
	address string
	client  *mongo.Client
	db      *mongo.Database
}

// Connect to db,
func Connect(ctx context.Context, address string) *Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	return &Database{
		context: ctx,
		address: address,
		client:  client,
	}
}
func Disconnect() {}
