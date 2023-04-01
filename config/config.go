package config

import (
	"strconv"

	"github.com/sjtommylee/go-dynamodb/utils/env"
)

type Config struct {
	Port        int
	Timeout     int
	Database    string
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "8080"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),
		Database:    env.GetEnv("Database", "sqlite3"),
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue))
	if err != nil {
		return 0
	}
	return num
}
