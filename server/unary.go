package main

import (
	"context"

	"github.com/Alex-911/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParams) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello World",
	}, nil
}
