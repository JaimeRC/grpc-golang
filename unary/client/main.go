package main

import (
	"context"
	pb "github.com/JaimeRC/grpc-golang/unary/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)


	// call Add service
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Pepe"})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)


	res, err = c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: "Pepe"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)
}
