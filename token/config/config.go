package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName   string `mapstructure:"SERVICE_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	RedisHostname string `mapstructure:"REDIS_HOSTNAME"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	APIKey        string `mapstructure:"API_KEY"`
	SecretKey     string `mapstructure:"SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	config = Config{
		ServiceName:   "token",
		ServerAddress: "0.0.0.0:8081",
		RedisHostname: "redis",
		RedisPort:     "6379",
		APIKey:        "SOME API KEY",
		SecretKey:     "SOME SECRET",
	}

	return config, nil
}

// compeltely useelss for some reason
func LoadViperConfig() (config Config, err error) {

	viper.SetDefault("SERVICE_NAME", "token")
	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:8081")
	viper.SetDefault("REDIS_HOSTNAME", "redis")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("API_KEY", "SOME API KEY")
	viper.SetDefault("SECRET_KEY", "SOME SECRET")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
