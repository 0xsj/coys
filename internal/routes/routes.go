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

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.SetConfigRouter()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) SetConfigRouter() {}

func RouterHealth() {}

func RouterProduct() {}

func EnableTimeout() {}

func EnableCors() {}

func EnableRecovery() {}

func EnableRequestID() {}

func EnableRealIP() {}
