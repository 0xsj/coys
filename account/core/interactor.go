/*
* abstracted layer way from repositories for service layer
 */
package core

import "context"

type UseCase interface {
	CreateAccount(ctx context.Context, email string, role Role) (*string, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

type UseCaseImpl struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &UseCaseImpl{repository: repository}
}

func (u *UseCaseImpl) CreateAccount(ctx context.Context, email string, role Role) (*string, error) {
	return u.repository.CreateAccount(ctx, email, role)
}

func (u *UseCaseImpl) GetAccountById(ctx context.Context, id string) (*Account, error) {
	return u.repository.GetAccountById(ctx, id)
}
