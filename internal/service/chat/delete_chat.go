package chat

import (
	"context"

	"github.com/bogdanove/chat-server/internal/repository/chat/converter"
)

const deleteAction = "DELETE"

// DeleteChat - удаление чата из системы по его идентификатору
func (c *chatService) DeleteChat(ctx context.Context, id int64) error {

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = c.chatRepository.DeleteChat(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = c.chatRepository.SaveLog(ctx, converter.FromServiceToLogRepo(id, deleteAction))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
