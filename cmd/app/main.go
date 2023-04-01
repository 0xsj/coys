package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sjtommylee/go-dynamodb/config"
	"github.com/sjtommylee/go-dynamodb/internal/repository/adapter"
	"github.com/sjtommylee/go-dynamodb/internal/repository/instance"
	"github.com/sjtommylee/go-dynamodb/internal/routes"
	"github.com/sjtommylee/go-dynamodb/utils/logger"
)

// new server struct for dependency injection.
type Server struct {
	Port       string
	Repository adapter.Interface
}

func NewServer(port string, repository adapter.Interface) *Server {
	return &Server{
		Port:       port,
		Repository: repository,
	}
}

func StartServer() {}

// main entry point
func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("service starting....", nil)
	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migrations", err)
		}
	}
	logger.PANIC("", CheckTables(connection))
	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running on port", port)
	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error

	return errors
}

func CheckTables(connection *dynamodb.DynamoDB) error {
	// response, err := conenction.ListTables()
	return nil
}
