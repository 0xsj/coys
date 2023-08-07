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
		ServerAddress: "127.0.0.1:8081",
		RedisHostname: "localhost",
		RedisPort:     "6379",
		APIKey:        "SOME API KEY",
		SecretKey:     "SOME SECRET",
	}

	return config, nil
}

// compeltely useless for some reason
func LoadViperConfig(name string) (config Config, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("env")
	viper.SetConfigName(name)
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
