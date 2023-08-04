package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	MongoHost       string `mapstructure:"MONGO_HOSTNAME"`
	MongoPort       string `mapstructure:"MONGO_PORT"`
	DatabaseName    string `mapstructure:"DATABASE_NAME"`
	CollectionItems string `mapstructure:"COLLECTION_ITEMS"`
	ApiKey          string `mapstructure:"API_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading configuration: %s", err)
		log.Fatal(err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshalling configuration: %s", err)
		log.Fatal(err)
	}
	return
}
