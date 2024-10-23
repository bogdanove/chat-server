package repository

import (
	"context"

	service "github.com/bogdanove/chat-server/internal/model"
	"github.com/bogdanove/chat-server/internal/repository/chat/model"
)

// ChatRepository - интерфейс функций репо слоя
type ChatRepository interface {
	CreateChat(ctx context.Context, req *service.Chat) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
	SaveLog(ctx context.Context, req *model.ChatLog) error
}
