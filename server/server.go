package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"grpcserver/chatimpl"

	"grpcchat/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chatimpl.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
