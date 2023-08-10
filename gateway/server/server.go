package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Address string
}

func (s *Server) Launch(bind func(*grpc.Server), opts ...grpc.ServerOption) error {
	server := grpc.NewServer(opts...)
	bind(server)

	reflection.Register(server)

	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Printf("Failed to start server on %s: %v", s.Address, err)
		return err
	}

	log.Printf("Server started on: %s", s.Address)
	if err = server.Serve(listener); err != nil {
		log.Printf("Error serving gRPC: %v", err)
		return err
	}

	return nil
}
