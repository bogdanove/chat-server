package chat

import (
	"github.com/bogdanove/chat-server/internal/client/db"
	"github.com/bogdanove/chat-server/internal/repository"
	"github.com/bogdanove/chat-server/internal/service"
)

type chatService struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewChatService - конструктор сервиса чат
func NewChatService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &chatService{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
