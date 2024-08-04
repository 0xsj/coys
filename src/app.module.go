package main

import (
	"net/http"

	"github.com/0xsj/coys/src/libs/di"
	"github.com/0xsj/coys/src/modules/auth"
)

// AppModule encapsulates all application modules.
type AppModule struct {
	authModule *auth.AuthModule
}

// NewAppModule creates a new AppModule with all necessary modules and container setup.
func NewAppModule() *AppModule {
	return &AppModule{
		authModule: auth.NewAuthModule(di.GlobalContainer),
	}
}

// RegisterRoutes registers routes for all modules.
func (m *AppModule) RegisterRoutes(router *http.ServeMux) {
	m.authModule.RegisterRoutes(router)
}

// RegisterControllers registers controllers for all modules.
func (m *AppModule) RegisterControllers() {
	m.authModule.RegisterControllers()
}

// RegisterProviders registers providers for all modules.
func (m *AppModule) RegisterProviders() {
	m.authModule.RegisterProviders()
}
