package main

import (
	"context"
	"log"
	"time"

	"grpcchat/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	message := &chat.Message{Body: "Hello From the Client!"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, message)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Body)
}
