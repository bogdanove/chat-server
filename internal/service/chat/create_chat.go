package chat

import (
	"context"

	"github.com/bogdanove/chat-server/internal/model"
	"github.com/bogdanove/chat-server/internal/repository/chat/converter"
)

const createAction = "CREATE"

// CreateChat - создания нового чата
func (c *chatService) CreateChat(ctx context.Context, req *model.Chat) (int64, error) {
	var id int64

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = c.chatRepository.CreateChat(ctx, converter.FromServiceToChatRepo(req))
		if errTx != nil {
			return errTx
		}

		errTx = c.chatRepository.SaveLog(ctx, converter.FromServiceToLogRepo(id, createAction))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
