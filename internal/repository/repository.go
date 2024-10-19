package repository

import (
	"context"

	"github.com/bogdanove/chat-server/internal/repository/chat/model"
)

// ChatRepository - интерфейс функций репо слоя
type ChatRepository interface {
	CreateChat(ctx context.Context, req *model.Chat) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
	SaveLog(ctx context.Context, req *model.ChatLog) error
}
