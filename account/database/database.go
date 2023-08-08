package database

import (
	"context"
	"fmt"
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

/*
@Connect
- takes in go's context, address as args. returns a database struct
- we attempt to connect to mongo passing in ctx, and the address, and store the variables in client, err
- if we do not get a connection, throw a fatal error
- else, log that there is a successful db connection
*/
func Connect(ctx context.Context, address string) *Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	fmt.Printf("Connected to database: %s\n", address)

	return &Database{
		context: ctx,
		address: address,
		client:  client,
	}
}

/*
@Disconnect
- takes in the datbase as arguments
- if we have an err during disconnecting, we throw a fatal error
- else, we return a auccess message with the db address
*/
func (d *Database) Disconenct() {
	if err := d.client.Disconnect(d.context); err != nil {
		log.Fatal("Could not disconnect from the database:", err)
	}
	fmt.Printf("Disconnected to database: %s\n", d.address)
}

/*
@Collection
- (d *Database) - a pointer to the instance of the db struct
- takes in db name as string, collection name as args
*/
func (d *Database) Collection(dbName string, collectionName string) *mongo.Collection {
	return d.client.Database(dbName).Collection(collectionName)
}
