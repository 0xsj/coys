package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{}).Handler(next)
}

func (c *Config) SetTimeout(timeSeconds int) *Config {
	c.timeout = time.Duration(timeSeconds) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
