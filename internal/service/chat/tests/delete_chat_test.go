package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bogdanove/platform_common/pkg/db"
	txMocks "github.com/bogdanove/platform_common/pkg/db/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/bogdanove/chat-server/internal/repository"
	repo "github.com/bogdanove/chat-server/internal/repository/chat/model"
	repoMocks "github.com/bogdanove/chat-server/internal/repository/mocks"
	"github.com/bogdanove/chat-server/internal/service/chat"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = int64(1)

		repoErr = fmt.Errorf("repository error")

		log = &repo.ChatLog{
			ChatID: id,
			Action: "DELETE",
		}
	)

	tests := []struct {
		name          string
		args          args
		err           error
		chatRepoMock  chatRepoMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "delete chat repository success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: nil,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteChatMock.Expect(ctx, id).Return(nil)
				mock.SaveLogMock.Expect(ctx, log).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "delete chat repository error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			err: repoErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteChatMock.Expect(ctx, id).Return(repoErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepoMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			service := chat.NewChatService(chatRepoMock, txManagerMock)

			err := service.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
