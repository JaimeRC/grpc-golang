package main

import (
	"context"
	"fmt"
	pb "github.com/JaimeRC/grpc-golang/server_streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewCountdownClient(conn)

	// timer to start countdown for 10 sec
	time := int32(10)

	// call Start Service
	stream, err := c.Start(context.Background(), &pb.CountdownRequest{Timer: time})
	if err != nil {
		log.Fatalf("failed to start timer: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}
		fmt.Println(msg)
	}

}
