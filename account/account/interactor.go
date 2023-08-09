/*
Responsible for handling our core logic
- encapsulating application behavior here
*/

package account

import "context"

// define contract for core application behavior
// specifies methods that represents use cases the application needs to support
type UseCase interface {
	CreateAccount(ctx context.Context, phoneNumber string, role Role) (*string, error)
	GetAccountById(ctx context.Context, id string) (*Account, error)
}

// implementation of the UseCase interface.
// holds a reference to the Repository that is used to interact with the data store
// struct encapsulates the actual business logic
type UseCaseImpl struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &UseCaseImpl{repository: repository}
}

func (u *UseCaseImpl) CreateAccount(ctx context.Context, phoneNumber string, role Role) (*string, error) {
	return u.repository.CreateAccount(ctx, phoneNumber, role)
}

func (u *UseCaseImpl) GetAccountById(ctx context.Context, id string) (*Account, error) {
	return u.repository.GetAccountById(ctx, id)
}
