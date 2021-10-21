package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

func (*UserService) AddUserVerbose(payload *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Initing...",
		User:   &pb.User{},
	})

	// aguardar 3 segundos
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

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	// Insertion in database stay here

	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatal("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.Id,
			Name:  req.Name,
			Email: req.Email,
		})

		fmt.Println("Adding", req.GetName())
	}

}
