package main

import (
	"fmt"
	pb "github.com/JaimeRC/grpc-golang/bi_directional_streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	port = ":50054"
)

type server struct {
	pb.UnimplementedFeedsServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not start the server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFeedsServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}

// Broadcast reads client stream and broadcasts received feeds
func (*server) Broadcast(stream pb.Feeds_BroadcastServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("could not recive from stream: %v", err)
			return err
		}
		feed := "New Feed Recieved: " + msg.GetFeed()
		fmt.Println("sending new feed...")
		stream.Send(&pb.FeedResponse{Feed: feed})
	}
}
