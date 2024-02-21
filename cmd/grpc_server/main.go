package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/dmtrybogdanov/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("Received CreateRequest: %v", req)

	return &chat_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*chat_v1.DeleteResponse, error) {
	log.Printf("Received DeleteRequest: %v", req)

	return &chat_v1.DeleteResponse{DeleteResponse: &empty.Empty{}}, nil

}

func (s *server) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*chat_v1.SendMessageResponse, error) {
	log.Printf("Received SendMessageRequest: %v", req)

	return &chat_v1.SendMessageResponse{SendMessageResponse: &empty.Empty{}}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
