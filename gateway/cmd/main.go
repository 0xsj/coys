package main

import (
	"context"
	"gateway/authentication"
	"gateway/config"
	pb "gateway/generated"
	"gateway/server"
	token "gateway/token"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		authenticationService := authentication.NewService(cfg.AuthenticationAddress)
		authenticationClient := authentication.NewClient(cfg.AuthenticationAddress)
		authInterceptor := server.NewInterceptor("Authorization", func(ctx context.Context, info *grpc.UnaryServerInfo, header string) error {
			if _, ok := info.Server.(authentication.ServiceImpl); ok {
				return nil
			}
			_, err := authenticationClient.VerifyAccess(ctx, &pb.VerifyAccessRequest{AccessToken: header})
			if err != nil {
				return status.Errorf(codes.Unauthenticated, "Invalid access token")
			}
			return nil
		})

		tokenService := token.NewService(cfg.TokenAddress)

		grpcServer := server.Server{Address: cfg.ServerAddress}

		grpcServer.Launch(func(server *grpc.Server) {
			pb.RegisterAuthenticationServiceServer(server, authenticationService)
		})

		grpcServer.Launch(func(server *grpc.Server) {
			pb.RegisterTokenServiceServer(server, tokenService)
		}, authInterceptor)

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
