package main

import (
	"context"
	"github.com/prasannakumar1989/go-grpc-example/chat"
	"log"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.Message) (*chat.Message, error) {
	log.Printf("Received message body from client: %v", message.Body)
	return &chat.Message{Body: "Hello From the server!"}, nil
}
