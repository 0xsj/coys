package main

import (
	"account/db"
	"context"
	"fmt"
)

var mongo_host = ""
var mongo_port = ""
var server_adddress = ""

func main() {
	db := db.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", mongo_host, mongo_port))
	defer db.Disconnect()
	// grpcServer := server.Server{Address: server_adddress}
	// grpcServer.Launch(func(server *grpc.Server) {
	// 	pb.RegisterAccountServiceServer(server, accountService)
	// }, authInterceptor)
}
