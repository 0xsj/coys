package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewInterceptor(name string, callback func(ctx context.Context, header string) error) grpc.ServerOption {
	// returns a grpc unary interceptor
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		return handler(ctx, req)
	})
}
