package main

import (
	"io"
	"log"

	pb "github.com/sneha31415/grpc/proto"
)

// since its a bidir stream, both client and server will have send and receive
func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if (err != nil) {
			return err
		}

		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}