package main

import (
	"log"
	"time"

	"github.com/Alex-911/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got Request with names %v", req.Names)
	for _, name := range req.Names {
		res := &proto.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
