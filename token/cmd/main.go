package main

import (
	"context"
	"fmt"
	"log"
	"token/config"
	"token/database"
)

func main() {
	fmt.Println("Main running")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := database.NewClient(context.Background(), fmt.Sprintf("%s:%s", cfg.RedisHostname, cfg.RedisPort))
	defer client.Close()

}
