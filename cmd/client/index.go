package main

import (
	"context"
	"log"

	"github.com/GiovannyLucas/grpc-golang/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	// when nobody are using this connection, it's closed
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	payload := &pb.User{
		Id:    "0",
		Name:  "Giovanny",
		Email: "giovanny@mail.com",
	}

	res, err := client.AddUser(context.Background(), payload)

	if err != nil {
		log.Fatalf("Could not to make gRPC request: %v", err)
	}

	log.Println(res)
}
