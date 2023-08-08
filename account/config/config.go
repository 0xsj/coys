package config

type Config struct {
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	MongoHostname   string `mapstructure:"MONGO_HOSTNAME"`
	MongoPort       string `mapstructure:"MONGO_PORT"`
	DatabaseName    string `mapstructure:"DATABASE_NAME"`
	CollectionItems string `mapstructure:"COLLECTION_ITEMS"`
	ApiKey          string `mapstructure:"API_KEY"`
}

func LoadConfig() (config Config, err error) {
	config = Config{
		ServiceName:     "account",
		ServerAddress:   "127.0.0.1:80802",
		MongoHostname:   "localhost",
		MongoPort:       "6379",
		DatabaseName:    "account",
		CollectionItems: "items",
		ApiKey:          "SOME KEY",
	}
	return config, nil
}
