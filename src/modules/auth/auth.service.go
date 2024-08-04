package auth

import "github.com/0xsj/coys/src/libs/di"

type AuthService struct {}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s*AuthService) Provide() interface{} {
	return NewAuthService()
}

func init() {
	di.RegisterInjectable(&AuthService{})
}