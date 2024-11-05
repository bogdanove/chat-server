package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/bogdanove/chat-server/internal/api/chat"
	"github.com/bogdanove/chat-server/internal/model"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/bogdanove/chat-server/internal/service"
	serviceMocks "github.com/bogdanove/chat-server/internal/service/mocks"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *chat_v1.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = int64(1)
		chatTitle = "some title"
		iDS       = []int64{1, 2, 3}

		serviceErr = fmt.Errorf("service error")

		req = &chat_v1.CreateRequest{
			Ids:       iDS,
			ChatTitle: chatTitle,
		}

		info = &model.Chat{
			IDs:       iDS,
			ChatTitle: chatTitle,
		}

		res = &chat_v1.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *chat_v1.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "create chat service success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "create chat service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, info).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			noteServiceMock := tt.chatServiceMock(mc)
			api := chat.NewServerImplementation(noteServiceMock)

			response, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}
