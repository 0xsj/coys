/*
Responsible for handling our core logic
- encapsulating application behavior here
*/

package account

type UseCase interface {
	CreateAccount()
}

type UseCaseImpl struct {
	repository Repository
}

func NewUseCase() {}

func (u *UseCaseImpl) CreateAccount() {}

// func GetAccountById() {}
