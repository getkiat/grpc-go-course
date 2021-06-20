package main

import (
	"context"
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

// SayHello implements helloworld.GreeterServer
func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet fn invoked with %v",in)
	firstName := in.GetGreeting().GetFirstName()
	lastName := in.GetGreeting().GetLastName()
	result := "Hello " + firstName + lastName
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
	// return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}


func main() {
	fmt.Println("Hello world, I'm server")

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
