/*
How we are going to interac with our database
This answers the question of how we are going to retrieve doamin objects / entities from our store
*/
package account

import (
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateAccount(ctx context.Context, phoneNumber string, role Role) (*string, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

type RepositoryImpl struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &RepositoryImpl{collection: collection}
}

// r *RepositoryImpl - method receiver. pointer to RepositoryIm
// we are returning a pointer to the string (the generated ID), and an error
func (r *RepositoryImpl) CreateAccount(ctx context.Context, phoneNumber string, role Role) (*string, error) {
	id := primitive.NewObjectID().Hex()
	if _, err := r.collection.InsertOne(ctx, Account{
		id,
		phoneNumber,
		role,
		Pending,
		time.Now().UnixMilli(),
	}); err != nil {
		return nil, err
	}
	return &id, nil
}

// Get Account, takes the context and the string id
func (r *RepositoryImpl) GetAccountById(ctx context.Context, id string) (*Account, error) {
	var result *Account
	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
