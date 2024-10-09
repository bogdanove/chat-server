package main

import (
	"context"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bogdanove/chat-server/internal/config"
	"github.com/bogdanove/chat-server/internal/config/env"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	listen, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

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

	_ = sq.InsertBuilder{}

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
