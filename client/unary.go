package main

import (
	"context"
	"log"
	"time"

	"github.com/Alex-911/go-grpc/proto"
)

func CallSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParams{})
	if err != nil {
		log.Fatalf("Could not get %v", err)
	}
	log.Printf("%s", res.Message)
}
