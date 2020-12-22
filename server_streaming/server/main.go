package main

import (
	pb "github.com/JaimeRC/grpc-golang/server_streaming/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedCountdownServer
}

const (
	port = ":50052"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn := grpc.NewServer()
	pb.RegisterCountdownServer(conn, &server{})

	if err := conn.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) Start(req *pb.CountdownRequest, stream pb.Countdown_StartServer) error {
	t := req.GetTimer()
	for t > 0 {
		res := pb.CountdownResponse{Count: t}
		_ = stream.Send(&res)
		t--
		time.Sleep(time.Second)
	}
	return nil
}
