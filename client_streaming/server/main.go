package main

import (
	pb "github.com/JaimeRC/grpc-golang/client_streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedSumAllServiceServer
}

const (
	port = ":50053"
)

func main()  {
	listener, err := net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSumAllServiceServer(s, &server{})

	if err:=s.Serve(listener); err != nil{
		log.Fatalf("failed to server: %v",err)
	}
}

func (*server) SumAll(stream pb.SumAllService_SumAllServer) error  {
	var sum int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.SumResponse{Result: sum})
		}

		if err != nil{
			log.Fatalf("could not recieve stream: %v", err)
		}

		sum = sum + msg.GetN()
	}
}