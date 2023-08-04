package core

import "context"

type UseCase interface {
	CreateAccount(ctx context.Context, email string, role Role) (*string, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

type UseCaseImpl struct {
	repository Repository
}
