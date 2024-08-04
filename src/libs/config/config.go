// config.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var AppConfig map[string]interface{}

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfigFileNotFound, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = make(map[string]interface{})
	if err := decoder.Decode(&AppConfig); err != nil {
		return fmt.Errorf("%w: %s", ErrConfigLoadFailed, err)
	}
	return nil
}

func GetConfig(key string) interface{} {
	return AppConfig[key]
}
