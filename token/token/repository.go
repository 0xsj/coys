package token

import (
	"context"
	"time"
	"token/config"

	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	GenerateTokenPair(ctx context.Context, id string) (*Pair, error)
}

type RepositoryImpl struct {
	config config.Config
	client *redis.Client
}

func NewRepository(config config.Config, client *redis.Client) Repository {
	return &RepositoryImpl{config: config, client: client}
}

/*
@GenerateTokenPair
- use theunix timestamp to claim in the jwt
- create a new claim based on the id, iat.
- sign the token with the secret key
- refresh should be the same process, but we set the expiratino to 30 days
*/
func (r *RepositoryImpl) GenerateTokenPair(ctx context.Context, id string) (*Pair, error) {
	timestamp := time.Now().Unix()
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": timestamp,
		"exp": timestamp + (time.Minute * 30).Milliseconds(),
	}).SignedString([]byte(r.config.SecretKey))
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": timestamp,
		"exp": timestamp + (time.Hour * 24 * 30).Milliseconds(),
	}).SignedString([]byte(r.config.SecretKey))
	if err != nil {
		return nil, err
	}
	return &Pair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
