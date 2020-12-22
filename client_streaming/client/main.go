package main

import (
	"context"
	"fmt"
	pb "github.com/JaimeRC/grpc-golang/client_streaming/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50053"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSumAllServiceClient(conn)

	stream, err := c.SumAll(context.Background())
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// send 0 to 10 numbers to the stream
	for i := 0; i <= 10; i++ {
		fmt.Printf("sending %v into the stream\n", i)
		stream.Send(&pb.NumberRequest{N: int32(i)})
		time.Sleep(100 * time.Microsecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recive response: %v", err)
	}

	fmt.Println("Sum of numbers: ", res.Result)
}
