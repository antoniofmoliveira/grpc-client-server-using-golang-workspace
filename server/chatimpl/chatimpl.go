package chatimpl

import (
	context "context"
	"grpcchat/chat"
	"log"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	log.Printf("Received: %v", message.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}
