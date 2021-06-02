package main

import (
	"context"
	pb "github.com/tcw/bazel-example/server-go/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (server) Echo(_ context.Context, in *pb.EchoIn) (*pb.EchoOut, error) {
	return &pb.EchoOut{
		Message: in.Message,
	}, nil
}

var _ = &server{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
