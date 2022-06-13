package main

import (
	"context"
	"fmt"
	"log"

	"github.com/clebersimm/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to grpc server: %v", err)
	}
	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "teste",
		Email: "teste",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make grpc request: %v", err)
	}
	fmt.Println(res)
}
