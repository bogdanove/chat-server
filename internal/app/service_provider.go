package app

import (
	"context"
	"log"

	"github.com/bogdanove/chat-server/internal/api/chat"
	server "github.com/bogdanove/chat-server/internal/api/chat"
	"github.com/bogdanove/chat-server/internal/client/db"
	"github.com/bogdanove/chat-server/internal/client/db/pg"
	"github.com/bogdanove/chat-server/internal/client/db/transaction"
	"github.com/bogdanove/chat-server/internal/closer"
	"github.com/bogdanove/chat-server/internal/config"
	"github.com/bogdanove/chat-server/internal/config/env"
	"github.com/bogdanove/chat-server/internal/repository"
	chatRepo "github.com/bogdanove/chat-server/internal/repository/chat"
	"github.com/bogdanove/chat-server/internal/service"
	chatService "github.com/bogdanove/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatServer *server.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig - конфигурация подключения к бд
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig - конфигурация сервера GRPC
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// DBClient - клиент для базы данных
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager - менеджер транзакций
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// ChatRepository - инициализация репозитория чата
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewChatRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

// ChatService - инициализация сервиса чата
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

// ChatImpl - инициализация сервера
func (s *serviceProvider) ChatImpl(ctx context.Context) *server.Server {
	if s.chatServer == nil {
		s.chatServer = chat.NewServerImplementation(s.ChatService(ctx))
	}

	return s.chatServer
}
