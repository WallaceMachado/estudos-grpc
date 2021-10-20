package services

import (
	"context"

	"github.com/wallacemachado/estudos-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	// se adcionar um seviço que não existe no protob não vai dar problema
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
