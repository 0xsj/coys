package config

type Config struct {
	Port        int
	Timeout     int
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        8008,
		Timeout:     30,
		DatabaseURI: "",
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	// num, err := strconv.Atoi(env.Geten)
	return 0
}
