package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

// define struct namd config, with a timeout field
type Config struct {
	timeout time.Duration
}

// function that returns an empty pointer to the Config struct we defined
func NewConfig() *Config {
	return &Config{}
}

// takes in HTTP handler as input, and returns a new http.handler with CORS enabled
// the resulting http handler, with the new Options struct, is returned  with input http.Handler passed as args
func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           5, // in seconds
	}).Handler(next)
}

// set the timeout field of the Config struct with a duration in seconds
// takes in int as input, then multplied by time.Second to get the duration in seconds
func (c *Config) SetTimeout(timeSeconds int) *Config {
	c.timeout = time.Duration(timeSeconds) * time.Second
	return c
}

// returns the timeout field as time.Duration
func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
