package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/getkiat/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "Meow"
)

func main() {
	fmt.Println("Hello I'm client")

	// grpc default with ssl, for demo non-ssl
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client: %f",c)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &pb.GreetRequest{
		Greeting: &pb.Greeting{
			FirstName: name,
			LastName: "tan",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResult())
}
