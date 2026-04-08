// implementation of server streaming
package main

import (
	"log"
	"time"

	pb "github.com/sneha31415/grpc/proto"
)

// unlike unary we don't have one response for a request to return.
// We don't return responses; we send them via the stream object using stream.Send
func(s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		// on sucess .Send() returns nil
		if err := stream.Send(res); err != nil {
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}