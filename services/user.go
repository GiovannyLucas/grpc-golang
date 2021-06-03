package services

import (
	"context"

	"github.com/GiovannyLucas/grpc-golang/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
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
