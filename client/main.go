package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcAssignment/user"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddUser(ctx, &pb.AddUserRequest{Name: "Zhartas", Email: "satibaevzhartas@gmail.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("id:%v\n name: %s\n email: %v\n", r.GetId(), r.GetName(), r.GetEmail())
	r, err = c.AddUser(ctx, &pb.AddUserRequest{Name: "Tima", Email: "tima@gmail.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("id:%v\n name: %s\n email: %v\n", r.GetId(), r.GetName(), r.GetEmail())

	byId, err := c.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Print(err)
	}
	fmt.Println(byId)

	listAll, err := c.ListAllUser(ctx, &pb.ListAllUserRequest{})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("all users %v\n", listAll)
}
