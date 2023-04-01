package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (s *Server) StartServer() {
	errors := Migrate(s.Repository)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migrations", err)
		}
	}

	logger.PANIC("", CheckTables(s.Repository))
	port := fmt.Sprintf(":%v", s.Port)
	router := routes.NewRouter().SetRouters(s.Repository)
	logger.INFO("service is running on port", port)
	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

// main entry point
func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("service starting....", nil)
	server := NewServer(strconv.Itoa(configs.Port), repository)
	server.StartServer()
}

func Migrate(connection adapter.Interface) []error {
	var errors []error
	CallMigrateAndAppendError(&errors, connection)
	return errors
}

func CallMigrateAndAppendError(errors *[]error, connection adapter.Interface) {

}

func CheckTables(connection adapter.Interface) (*dynamodb.ListTablesOutput, error) {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return err
}
