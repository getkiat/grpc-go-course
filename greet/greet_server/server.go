package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/getkiat/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

// Default port for grpc
const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err:= s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
