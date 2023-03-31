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

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.newHandler(repository)
	r.router.Route("/health", func(route chi.Router) {
		route.Post
		route.Get
		route.Put
		route.Delete
		route.Options
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.NewHandler(repository)
	r.router.Route("/product", func(route, chi.Router) {
		route.Post
		route.Get
		route.Put
		route.Delete
		route.Options
	})
}

func (r *Router) EnableTimeout() *Router {
	return r
}

func (r *Router) EnableLogger() *Router {
	return r
}

func (r *Router) EnableCors() *Router {
	return r
}

func (r *Router) EnableRecovery() *Router {
	return r
}

func (r *Router) EnableRequestID() *Router {
	return r
}

func (r *Router) EnableRealIP() *Router {
	return r
}
