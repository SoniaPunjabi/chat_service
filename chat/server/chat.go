package main

import (
	"bufio"
	pb "chat_service/chat/proto"
	"fmt"
	"io"
	"log"
	"os"
)

//var resp = make(chan string)

func (s *Server) Chat(stream pb.ChatService_ChatServer) error {

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		fmt.Printf("Allen-> %s", req.Message)
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		go func() {
			for {
				inputReader := bufio.NewReader(os.Stdin)
				input, _ := inputReader.ReadString('\n')

				err := stream.Send(&pb.ChatResponse{
					Response: input,
				})
				if err != nil {
					log.Fatalf("error while sending data to client stream: %v\n", err)
				}
			}
		}()
	}

}
