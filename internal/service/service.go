package service

import (
	"context"

	"github.com/bogdanove/chat-server/internal/model"
)

// ChatService - интерфейс функций сервисного слоя
type ChatService interface {
	CreateChat(ctx context.Context, req *model.Chat) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}
