package config

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
	return
}
