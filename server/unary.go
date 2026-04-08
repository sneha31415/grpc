package main

import (
	"context"

	pb "github.com/sneha31415/grpc/proto"
)

// (s *helloServer) means: “this function is a method of helloServer.”
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, This is the output of a unary rpc call",
	}, nil
}