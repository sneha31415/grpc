package main

import (
	"log"
	"net"

	pb "github.com/sneha31415/grpc/proto"
	"google.golang.org/grpc"
)

// port on which server will run
const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}
func main(){
	// listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start : %v", err)
	}
}