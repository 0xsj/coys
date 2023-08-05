package account

import (
	pb "gateway/generated"
	"log"

	"google.golang.org/grpc"
)

func NewClient(address string) pb.AuthenticationServiceClient {
	connection, err := grpc.Dial(address)

	if err != nil {
		log.Fatal(err)
	}
	return pb.NewAuthenticationServiceClient(connection)

}
