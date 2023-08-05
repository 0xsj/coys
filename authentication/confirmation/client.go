package confirmation

import (
	pb "authentication/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewClient(address string) pb.ConfirmationServiceClient {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewConfirmationServiceClient(connection)
}
