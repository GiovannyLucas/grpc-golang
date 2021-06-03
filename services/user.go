package services

import (
	"context"
	"time"

	"github.com/GiovannyLucas/grpc-golang/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, payload *pb.User) (*pb.User, error) {
	// Insertion in database stay here

	return &pb.User{
		Id:    "123",
		Name:  payload.GetName(),
		Email: payload.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(payload *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Initing...",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting on database...",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted!",
		User: &pb.User{
			Id:    "123",
			Name:  payload.GetName(),
			Email: payload.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  payload.GetName(),
			Email: payload.GetEmail(),
		},
	})

	return nil
}
