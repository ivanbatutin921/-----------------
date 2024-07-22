package main

import (
	"fmt"
	"net"

	db "root/mk/internal/database"
	pb "root/mk/internal/proto"
	"root/mk/internal/server"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()
	//db.Migrate()


	lis, err := net.Listen("tcp", ":50050")
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
