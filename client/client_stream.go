package main

import (
	"context"
	"log"
	"time"

	"github.com/Alex-911/go-grpc/proto"
)

func CallSayHelloClientStreaming(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Client stream started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send name: %v", err)
	}
	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending : %v", err)
		}
		log.Printf("Sen the request with name %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Stream Finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("Got Response with name : %v", res.Messages)
}
