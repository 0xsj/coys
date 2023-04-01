package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sjtommylee/go-dynamodb/internal/repository/adapter"
	// HealthHandler "github.com/sjtommylee/internal/handlers/health"
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
	// r.RouterHealth(repository)
	// r.RouterProduct(repository)

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

// func (r *Router) RouterHealth(repository adapter.Interface) {
// 	handler := HealthHandler.newHandler(repository)
// 	r.router.Route("/health", func(route chi.Router) {
// 		route.Post("/", handler.Post)
// 		route.Get("/", handler.Get)
// 		route.Put("/", handler.Put)
// 		route.Delete("/", handler.Delete)
// 		route.Options("/", handler.Options)
// 	})
// }

// func (r *Router) RouterProduct(repository adapter.Interface) {
// 	handler := ProductHandler.NewHandler(repository)
// 	r.router.Route("/product", func(route, chi.Router) {
// 		route.Post("/", handler.Post)
// 		route.Get("/", handler.Get)
// 		route.Put("/{ID}", handler.Put)
// 		route.Delete("/{ID}", handler.Delete)
// 		route.Options("/", handler.Options)
// 	})
// }

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCors() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecovery() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
