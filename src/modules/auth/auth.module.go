package auth

import "net/http"


type AuthModule struct {}

func NewAuthModule() *AuthModule {
	return &AuthModule{}
}

func (m *AuthModule) RegisterRoutes(router *http.ServeMux) {}

func (m* AuthModule) RegisterControllers() {}

func (m *AuthModule) RegisterProviders() {}

