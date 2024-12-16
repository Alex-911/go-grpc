package main

import (
	"log"

	"github.com/Alex-911/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client did not connect %v", err)
	}
	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	CallSayHello(client)

	names := &proto.NamesList{
		Names: []string{"Rohit", "Aarti", "Motki"},
	}

	CallSayHelloServerStreaming(client, names)

}
