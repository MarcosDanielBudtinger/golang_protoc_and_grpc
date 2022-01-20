package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@J.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not MAKE gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@J.com",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not MAKE gRPC request: %v", err)
	}
	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Could not receive the msg %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		&pb.User{
			Id:    "m1",
			Name:  "Marcos",
			Email: "m@m.com",
		},
		&pb.User{
			Id:    "m2",
			Name:  "Marcos 2",
			Email: "m@m2.com",
		},
		&pb.User{
			Id:    "m3",
			Name:  "Marcos 3",
			Email: "m@m3.com",
		},
		&pb.User{
			Id:    "m4",
			Name:  "Marcos 4",
			Email: "m@m4.com",
		},
		&pb.User{
			Id:    "m5",
			Name:  "Marcos 5",
			Email: "m@m5.com",
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
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating response: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "m1",
			Name:  "Marcos",
			Email: "m@m.com",
		},
		&pb.User{
			Id:    "m2",
			Name:  "Marcos 2",
			Email: "m@m2.com",
		},
		&pb.User{
			Id:    "m3",
			Name:  "Marcos 3",
			Email: "m@m3.com",
		},
		&pb.User{
			Id:    "m4",
			Name:  "Marcos 4",
			Email: "m@m4.com",
		},
		&pb.User{
			Id:    "m5",
			Name:  "Marcos 5",
			Email: "m@m5.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
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
				log.Fatalf("Error receiving data: %v", err)
			}
			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
