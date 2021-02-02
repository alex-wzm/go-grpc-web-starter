package main

import (
	"fmt"
	"log"
	"net"
	"server/echo"

	"google.golang.org/grpc"
)

var port = "9090"

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := echo.Server{}

	grpcServer := grpc.NewServer()

	echo.RegisterEchoServiceServer(grpcServer, &s)

	log.Printf("Servering Echo on port %s...", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server on port %s: %v", port, err)
	}
}
