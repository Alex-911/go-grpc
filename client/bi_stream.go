package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/Alex-911/go-grpc/proto"
)

func CallSayHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Bidirectional Stream Started")
	stream, err := client.SayHelloBidirectionStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not sent name %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming Completed")
}
