package core

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateAccount(ctx context.Context, email string, role Role) (*string, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

type RepositoryImpl struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &RepositoryImpl{collection: collection}
}

func (r *RepositoryImpl) CreateAccount(ctx context.Context, email string, role Role) (*string, error) {
	id := primitive.NewObjectID().Hex()
	if _, err := r.collection.InsertOne(ctx, Account{
		id,
		email,
		role,
		PendingConfirmation,
		time.Now().UnixMilli(),
	}); err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *RepositoryImpl) GetAccountById(ctx context.Context, id string) (*Account, error) {
	var result *Account
	if err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
