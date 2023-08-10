package config

type Config struct {
	ServiceName         string `mapstructure:"SERVICE_NAME"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	RedisHostname       string `mapstructure:"REDIS_HOSTNAME"`
	RedisPort           string `mapstructure:"REDIS_PORT"`
	MongoHost           string `mapstructure:"MONGO_HOST"`
	MongoPort           string `mapstructure:"MONGO_PORT"`
	DatabaseName        string `mapstructure:"DATABASE_NAME"`
	CollectionItems     string `mapstructure:"COLLECTION_ITEMS"`
	TokenAddress        string `mapstructure:"TOKEN_ADDRESS"`
	AccountAddress      string `mapstructure:"ACCOUNT_ADDRESS"`
	ConfirmationAddress string `mapstructure:"CONFIRMATION_ADDRESS"`
	ApiKey              string `mapstructure:"API_KEY"`
}

func LoadConfig() (config Config, err error) {
	config = Config{
		ServiceName:         "authentication",
		ServerAddress:       "0.0.0.0:8080",
		RedisHostname:       "0.0.0.0",
		RedisPort:           "6379",
		MongoHost:           "0.0.0.0",
		MongoPort:           "27017",
		DatabaseName:        "authentication",
		CollectionItems:     "items",
		TokenAddress:        "0.0.0.0:8081",
		AccountAddress:      "0.0.0.0:8082",
		ConfirmationAddress: "0.0.0.0:8083",
		ApiKey:              "nada",
	}
	return config, nil
}
