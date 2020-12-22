package main

import (
	"context"
	pb "github.com/JaimeRC/grpc-golang/unary/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{
	pb.UnimplementedGreeterServer
}

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	conn := grpc.NewServer()
	pb.RegisterGreeterServer(conn, &server{})

	if err := conn.Serve(listener); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}
