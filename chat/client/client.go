package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/SoniaPunjabi/chat_service/chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "server:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect : %v\n", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	doChat(c)
}
func doChat(c pb.ChatServiceClient) {

	for {
		stream, err := c.Chat(context.Background())
		if err != nil {
			log.Fatalf("Error while creating stream: %v\n", err)
		}

		req := &pb.ChatRequest{}

		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		req.Message = input

		stream.Send(req)

		go func() {

			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Printf("Error while receiving: %v\n", err)
					break
				}
				fmt.Printf("Bob-> %v", res.Response)
			}

		}()

	}
}
