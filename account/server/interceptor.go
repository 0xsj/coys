/*
provides a gRPC server interceptor that intercepts incoming RPC requests.
This interceptor allows interaction with the proto message or context before or after it is sent by the client or server.
*/
package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NewInterceptor creates a new gRPC server interceptor.
func NewInterceptor(name string, callback func(ctx context.Context, header string) error) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Missing or invalid metadata: %s", name)
		}

		// Extract header
		header := md.Get(name)
		if len(header) < 1 {
			return nil, status.Errorf(codes.InvalidArgument, "Missing or invalid header")
		}

		// Callback to process header
		if err := callback(ctx, header[0]); err != nil {
			return nil, status.Errorf(codes.Internal, "Error processing header: %v", err)
		}

		// Handle the request and propagate any errors
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	})
}
