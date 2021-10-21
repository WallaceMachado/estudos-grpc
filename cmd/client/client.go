package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wallacemachado/estudos-grpc/pb"
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

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)

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

		log.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wallace 1",
			Email: "w1@b.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wallace 2",
			Email: "w2@b.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wallace 3",
			Email: "w3@b.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wallace 4",
			Email: "w4@b.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving request: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wallace 1",
			Email: "w1@b.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wallace 2",
			Email: "w2@b.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wallace 3",
			Email: "w3@b.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wallace 4",
			Email: "w4@b.com",
		},
	}

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Could not receive datas: %v", stream)
				break
			}

			log.Printf("Receiving user %v with status : %v\n", res.User.GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
