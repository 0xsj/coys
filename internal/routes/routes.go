package routes

import (
	"github.com/go-chi/chi"
	"github.com/sjtommylee/go-dynamodb/internal/repository/adapter"
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

func (r *Router) SetConfigRouter() {
	r.EnableCors()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecovery()
	r.EnableRequestID()
	r.EnableRealIP()

}

func (r *Router) RouterHealth(repository adapter.Interface) {}

func (r *Router) RouterProduct() {}

func EnableTimeout() {}

func EnableCors() {}

func EnableRecovery() {}

func EnableRequestID() {}

func EnableRealIP() {}
