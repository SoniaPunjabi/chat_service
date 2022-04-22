package main

import (
	"bufio"
	pb "chat_service/chat/proto"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

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
