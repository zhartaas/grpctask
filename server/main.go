package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcAssignment/user"
	"log"
	"net"
)

var Users []*pb.User
var lastUserId int32 = 0

var (
	port = flag.Int("port", 50051, "Server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.User, error) {
	user := &pb.User{Id: lastUserId + 1, Name: in.GetName(), Email: in.GetEmail()}
	Users = append(Users, user)
	lastUserId++
	return user, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	for _, user := range Users {
		if user.Id == in.Id {
			return user, nil
		}
	}
	return nil, errors.New("not found")
}

func (s *server) ListAllUser(ctx context.Context, in *pb.ListAllUserRequest) (*pb.ListOfUsers, error) {
	return &pb.ListOfUsers{Users: Users}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
