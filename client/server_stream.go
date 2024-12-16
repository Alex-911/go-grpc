package main

import (
	"io"
	"log"

	"github.com/Alex-911/go-grpc/proto"
	"golang.org/x/net/context"
)

func CallSayHelloServerStreaming(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
