package token

import "context"

type UseCase interface {
	GenerateTokenPair(ctx context.Context, payload string) (*Pair, error)
}

type UseCaseImpl struct {
	repository Repository
}

func NewUseCase(repository Repository) UseCase {
	return &UseCaseImpl{repository: repository}
}

func (u *UseCaseImpl) GenerateTokenPair(ctx context.Context, payload string) (*Pair, error) {
	return u.repository.GenerateTokenPair(ctx, payload)
}
