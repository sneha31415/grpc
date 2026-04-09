package main

import (
	"log"

	pb "github.com/sneha31415/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	// close the connection between client and server
	defer conn.Close()
	
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names:[]string{"Sneha", "Shubham", "Mummy", "Papa"},
	}

	//--- unary rpc call ---
	// callSayHello(client)

	// --- Server Streaming RPC ---
	// callSayHelloServerStream(client, names)

	// client streaming rpc
	callSayHelloClientStreaming(client, names)


}


// NewClient → “I have a road”
// NewGreetServiceClient → “I have a car that knows the route names (RPC methods) on that road”