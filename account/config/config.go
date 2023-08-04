package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName     string
	ServerAddress   string
	MongoHost       string
	MongoPort       string
	DatabaseName    string
	CollectionItems string
	ApiKey          string
}

func LoadConfig(name string) (config Config, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("env")
	viper.SetConfigName(name)

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
