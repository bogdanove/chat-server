package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bogdanove/platform_common/pkg/db"
	txMocks "github.com/bogdanove/platform_common/pkg/db/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/bogdanove/chat-server/internal/model"
	"github.com/bogdanove/chat-server/internal/repository"
	repo "github.com/bogdanove/chat-server/internal/repository/chat/model"
	repoMocks "github.com/bogdanove/chat-server/internal/repository/mocks"
	"github.com/bogdanove/chat-server/internal/service/chat"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(1)
		chatTitle = "some title"
		iDS       = []int64{1, 2, 3}

		repoErr = fmt.Errorf("repository error")

		req = &model.Chat{
			IDs:       iDS,
			ChatTitle: chatTitle,
		}
		log = &repo.ChatLog{
			ChatID: id,
			Action: "CREATE",
		}
	)

	tests := []struct {
		name          string
		args          args
		want          int64
		err           error
		chatRepoMock  chatRepoMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "create chat repository success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(id, nil)
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
			name: "create chat repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, repoErr)
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

			response, err := service.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}
