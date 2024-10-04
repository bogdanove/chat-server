package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bogdanove/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("create new chat with usernames: %v", req.Usernames)
	_ = ctx
	return &chat_v1.CreateResponse{
		Id: int64(len(req.Usernames)),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("delete chat for id: %d", req.Id)
	_ = ctx
	return new(emptypb.Empty), nil
}

func (s *server) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("incoming message received: %+v", req.Message)
	_ = ctx
	return new(emptypb.Empty), nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at: %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
