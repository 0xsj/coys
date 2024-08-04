package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

// App represents the application with a router, modules, and metadata.
type App struct {
	Router   *http.ServeMux
	Modules  []Module
	Metadata map[string]interface{}
}

// NewApp initializes a new App with given modules and sets up default values.
func NewApp(modules ...Module) *App {
	app := &App{
		Router:   http.NewServeMux(),
		Modules:  []Module{},
		Metadata: make(map[string]interface{}),
	}
	app.RegisterModules(modules...)
	return app
}

// RegisterModules iterates over modules and calls their registration functions.
func (a *App) RegisterModules(modules ...Module) {
	for _, module := range modules {
		if module == nil {
			continue
		}
		module.RegisterRoutes(a.Router)
		module.RegisterControllers()
		module.RegisterProviders()
		a.Modules = append(a.Modules, module)
	}
}

// SetMetadata sets a value in the metadata map.
func (a *App) SetMetadata(key string, value interface{}) {
	a.Metadata[key] = value
}

// GetMetadata retrieves a value from the metadata map.
func (a *App) GetMetadata(key string) interface{} {
	return a.Metadata[key]
}

// GetMetadataKeys returns a list of all metadata keys.
func (a *App) GetMetadataKeys() []string {
	var keys []string
	val := reflect.ValueOf(a.Metadata)
	for _, key := range val.MapKeys() {
		keys = append(keys, key.String())
	}
	return keys
}

// Start starts the HTTP server and gracefully handles shutdowns.
func (a *App) Start(port int) {
	a.SetMetadata("port", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.Router,
	}

	// Start server in a goroutine.
	go func() {
		fmt.Printf("Starting server on port %d\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")

	// Create a context with a timeout for shutting down.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}
	fmt.Println("Server gracefully stopped")
}
