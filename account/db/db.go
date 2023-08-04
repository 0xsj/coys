package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	context context.Context
	address string
	client  *mongo.Client
	db      *mongo.Database
}

func Connect(ctx context.Context, address string) *DB {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("connected to the database: %s\n", address)
	return &DB{
		context: ctx,
		address: address,
		client:  client,
	}
}

func (d *DB) Disconnect() {
	if err := d.client.Disconnect(d.context); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Disconnected from the database: %s\n", d.address)
}

func (d *DB) Collection(dbName string, collectionName string) *mongo.Collection {
	return d.client.Database(dbName).Collection(collectionName)
}
