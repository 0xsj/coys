package routes

import (
	"github.com/go-chi/chi"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig(),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters() *chi.Mux {}

func (r *Router) SetConfigRouter() {}

func RouterHealth() {}

func RouterProduct() {}

func EnableTimeout() {}

func EnableCors() {}

func EnableRecovery() {}

func EnableRequestID() {}

func EnableRealIP() {}
