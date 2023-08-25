package authentication

import (
	"authentication/account"
	"authentication/confirmation"
	"authentication/token"
	"context"
)

type UseCase interface {
	SignInByPhoneNumber(ctx context.Context, phoneNumber string) (*int64, error)
}

type UseCaseImpl struct {
	accountRepository      account.Repository
	confirmationRepository confirmation.Repository
	tokenRepository        token.Repository
}

func NewUseCase(accountRepository account.Repository, confirmationRepository confirmation.Repository, tokenRepository token.Repository) UseCase {
	return &UseCaseImpl{accountRepository: accountRepository, confirmationRepository: confirmationRepository, tokenRepository: tokenRepository}
}

func (u *UseCaseImpl) SignInByPhoneNumber(ctx context.Context, phoneNumber string) (*int64, error) {
	acc, _ := u.accountRepository.GetAccountByPhoneNumber(ctx, phoneNumber)
	if acc == nil {
		if _, err := u.accountRepository.CreateAccount(ctx, phoneNumber, account.User); err != nil {
			return nil, err
		}
	}
	return u.confirmationRepository.SendPhoneNumberConfirmation(ctx, phoneNumber)
}

// func (u *UseCaseImpl) SignOut(ctx context.Context, tokens *token.Pair) (*string, error) {
// 	access, refresh := tokens.AccessToken, tokens.RefreshToken
// 	return u.tokenRepository.
// }
