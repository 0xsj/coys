// loader.go
package config

import (
	"log"
)

func Init() {
	if err := LoadConfig("config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
