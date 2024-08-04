package app

import (
	"fmt"
	"net/http"
)

type App struct {
	Router *http.ServeMux
	Modules []Module
}

func NewApp() *App {
	return &App{
		Router: http.NewServeMux(),
		Modules: []Module{},
	}
}

func (a *App) RegisterRoutes(modules ...Module) {
	for _, module := range modules {
		module.RegisterRoutes(a.Router)
		a.Modules = append(a.Modules, module)
	}
}

func (a *App) Start(port interface{}) {
	fmt.Printf("starting server on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router); err != nil {
		panic(err)
	}
}