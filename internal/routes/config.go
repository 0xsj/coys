package routes

import (
	"net/http"
	"time"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {}

func (c *Config) SetTimeout(timeSeconds int) *Config {
	c.timeout = time.Duration(timeSeconds) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
