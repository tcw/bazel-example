package main

import (
	"context"
	"log"
	"time"
	pb "github.com/tcw/bazel-example/go/grpc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoIn{Message: "hello world"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo: %s", r.GetMessage())
}