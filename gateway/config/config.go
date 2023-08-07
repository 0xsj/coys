package config

type Config struct {
	ServerAddress         string `mapstructure:"SERVER_ADDRESS"`
	AuthenticationAddress string `mapstructure:"AUTHENTICATION_ADDRESS"`
	TokenAddress          string `mapstructure:"TOKEN_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	config = Config{
		ServerAddress:         "0.0.0.0:8000",
		AuthenticationAddress: "0.0.0.0:8080",
		TokenAddress:          "0.0.0.0:8090",
	}

	return config, nil
}
