package di

import (
	"fmt"
	"reflect"
	"sync"
)

// Container is a dependency injection container.
type Container struct {
	services map[string]reflect.Value
	mu       sync.RWMutex
}

// NewContainer creates a new instance of Container.
func NewContainer() *Container {
	return &Container{
		services: make(map[string]reflect.Value),
	}
}

// RegisterProvider registers a service provider with a given name.
func (c *Container) RegisterProvider(name string, provider interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.services[name] = reflect.ValueOf(provider)
}

// Resolve retrieves a service from the container.
func (c *Container) Resolve(name string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	service, found := c.services[name]
	if !found {
		return nil, fmt.Errorf("service %s not found", name)
	}

	return service.Interface(), nil
}


var GlobalContainer = NewContainer()