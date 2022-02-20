package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
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
	fmt.Println("Added user without stream, only a request and response")
	fmt.Println("------------------------------------------------------")
}

func AddUserVerbose(client pb.UserServiceClient) {
	fmt.Println("To add user with data stream")

	payload := &pb.User{
		Id:    "0",
		Name:  "Giovanny",
		Email: "giovanny@mail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), payload)

	if err != nil {
		log.Fatalf("Could not to make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive datas: %v", stream)
		}

		log.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Gio",
			Email: "gio@mail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Gio 2",
			Email: "gio2@mail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Gio 3",
			Email: "gio3@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error while sending message: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receive the response from server: %v", err)
	}

	fmt.Println(res)
}
