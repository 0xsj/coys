package authentication

import (
	"authentication/account"
	"authentication/confirmation"
	"authentication/token"
)

type UseCase interface {
}

type UseCaseImpl struct {
	accountRepository      account.Repository
	confirmationRepository confirmation.Repository
	tokenRepository        token.Repository
}

func NewUseCase(accountRepository account.Repository, confirmationRepository confirmation.Repository, tokenRepository token.Repository) UseCase {
	return &UseCaseImpl{accountRepository: accountRepository, confirmationRepository: confirmationRepository, tokenRepository: tokenRepository}
}
