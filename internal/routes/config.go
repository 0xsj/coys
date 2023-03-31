package routes

import "time"

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func Cors() {}

func SetTimeout() {}

func GetTimeout() {}
