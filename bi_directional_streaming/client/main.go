package main

import (
	"context"
	"fmt"
	pb "github.com/JaimeRC/grpc-golang/bi_directional_streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

const (
	address = "localhost:50054"
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v",err)
	}

	c := pb.NewFeedsClient(conn)
	// get client stream
	stream, err := c.Broadcast(context.Background())
	if err != nil{
		log.Fatalf("failed to call Broadcast: %v",err)
	}

	// make blocking channel
	waitc := make(chan struct{})

	//send feeds to the stream (go routine)
	go func() {
		for i := 1; i<=5;i++{
			feed:= "This is feed number " + strconv.Itoa(i)
			if err := stream.Send(&pb.FeedRequest{Feed:feed});err != nil{
				log.Fatalf("error while sending feed: %v",err)
			}
			time.Sleep(time.Second)
		}
		if err:= stream.CloseSend();err!= nil{
			log.Fatalf("failed to close stream: %v",err)
		}
	}()

	// receive feeds from the streaam (go routine)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF{
				close(waitc)
				return
			}
			if err != nil{
				log.Fatalf("failed to recieve: %v",err)
				close(waitc)
				return
			}
			fmt.Println("Bew feed received: ", msg.GetFeed())
		}
	}()
	<-waitc
}
