package main

import (
	"net/http"

	"github.com/0xsj/coys/src/modules/auth"
)

// IE: import gitub.com/0xsj/coys/src/modules/auth auth.module.go
// IE: import gitub.com/0xsj/coys/src/modules/profile profile.module.go

type AppModule struct {
	authModule *auth.AuthModule
}

func NewAppModule() *AppModule {
	appModule := &AppModule{
		authModule: auth.NewAuthModule(),
	}

	return appModule
}

func (m *AppModule) RegisterRoutes(router *http.ServeMux) {
	m.authModule.RegisterRoutes(router)
}
func (m *AppModule) RegisterControllers() {
	m.authModule.RegisterControllers()
}
func (m *AppModule) RegisterProviders() {
	m.authModule.RegisterProviders()
}