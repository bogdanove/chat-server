package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	chat_v1.UnimplementedChatV1Server
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

// Create - создания нового чата
func (s *server) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	log.Printf("create new chat with title: %v", req.ChatTitle)

	return &chat_v1.CreateResponse{
		Id: int64(len(req.Ids)),
	}, nil
}

// Delete - удаление чата из системы по его идентификатору
func (s *server) Delete(_ context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("delete chat for id: %d", req.Id)

	return &emptypb.Empty{}, nil
}

// SendMessage - отправка сообщения на сервер
func (s *server) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("incoming message received: %+v", req.Message)

	return &emptypb.Empty{}, nil
}
