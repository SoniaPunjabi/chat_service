package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/SoniaPunjabi/chat_service/chat/proto"

	"google.golang.org/grpc"
)

var addr string = ":50051"

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
