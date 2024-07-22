package main

import (
	"fmt"
	"net"

	
	pb "root/second_microservice/internal/proto"
	"root/second_microservice/internal/server"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	} else {
		fmt.Println("server started")
	}
	mk := grpc.NewServer()
	pb.RegisterUserServiceServer(mk, &server.Server{}) 
	if err := mk.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
