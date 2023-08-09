package main

import (
	"account/account"
	"account/config"
	"account/database"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", cfg.MongoHostname, cfg.MongoPort))
	defer db.Disconenct()

	accountRepository := account.NewRepository(db.Collection(cfg.DatabaseName, cfg.CollectionItems)
	accountUseCase := account.NewUseCase(accountRepository)

}
