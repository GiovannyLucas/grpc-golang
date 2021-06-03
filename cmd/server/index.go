package main

import (
	"log"
	"net"

	"github.com/GiovannyLucas/grpc-golang/pb"
	"github.com/GiovannyLucas/grpc-golang/services"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// instancing a new server gRPC
	grpcServer := grpc.NewServer()

	// registering new services
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	// serving the server on port specified
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
