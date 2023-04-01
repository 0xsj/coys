package main

import (
	"github.com/sjtommylee/go-dynamodb/config"
	"github.com/sjtommylee/go-dynamodb/internal/repository/instance"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
}

func Migrate() {}

func checkTables() {}
