package main

import (
	pb "chat_service/chat/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.ChatServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("listening on %s", addr)
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
