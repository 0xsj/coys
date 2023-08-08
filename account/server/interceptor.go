package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
@NewInterceptor
- creates a grpc interceptor function
- Interceptors in the context of gRPC play a crucial role in extending, modifying, or adding functionality to the communication process between gRPC clients and servers. Interceptors provide a way to intercept and process various aspects of gRPC requests and responses, allowing you to add cross-cutting concerns without modifying each individual service method.
*/

func NewInterceptor(name string, callback func(ctx context.Context, header string) error) grpc.ServerOption {
	// returns a grpc unary interceptor, anonymous function
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Error reading metadata")
		}

		header := md.Get(name)
		if len(header) < 1 {
			return nil, status.Errorf(codes.InvalidArgument, "Error reading header")
		}

		if err := callback(ctx, header[0]); err != nil {
			return nil, status.Errorf(codes.Internal, "Error processing header")
		}

		return handler(ctx, req)
	})
}
