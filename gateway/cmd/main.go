package main

import (
	"gateway/config"
	pb "gateway/generated"
	"gateway/server"
	token "gateway/token"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Launch the gRPC server in a goroutine
	go func() {
		grpcServer := server.Server{Address: cfg.ServerAddress}
		tokenService := token.NewService(cfg.TokenAddress)
		grpcServer.Launch(func(server *grpc.Server) {
			pb.RegisterTokenServiceServer(server, tokenService)
			reflection.Register(server)
		})
	}()

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define an HTTP route using Mux
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP via Mux!"))
	})

	// Create an HTTP server with Mux and listen on a port
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("HTTP server started on: %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
