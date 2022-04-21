package main

import (
	"log"

	pb "chat_service/chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect : %v\n", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	doChat(c)
}
