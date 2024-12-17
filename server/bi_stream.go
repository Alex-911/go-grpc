package main

import (
	"io"
	"log"

	"github.com/Alex-911/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionStreaming(stream proto.GreetService_SayHelloBidirectionStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &proto.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
