// auth/module.go
package auth

import (
	"net/http"

	"github.com/0xsj/coys/src/libs/di"
)

type AuthModule struct {
	container *di.Container
}

func NewAuthModule(container *di.Container) *AuthModule {
	return &AuthModule{container: container}
}

func (m *AuthModule) RegisterRoutes(router *http.ServeMux) {
	// Register your routes here, e.g., router.HandleFunc("/auth/login", m.LoginHandler)
}

// RegisterControllers registers controllers for the AuthModule.
func (m *AuthModule) RegisterControllers() {
	// Register controllers if needed
}

// RegisterProviders registers providers for services in the AuthModule.
func (m *AuthModule) RegisterProviders() {
}
