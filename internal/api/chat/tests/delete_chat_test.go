package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bogdanove/chat-server/internal/api/chat"
	"github.com/bogdanove/chat-server/internal/service"
	serviceMocks "github.com/bogdanove/chat-server/internal/service/mocks"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *chat_v1.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = int64(1)

		serviceErr = fmt.Errorf("service error")

		req = &chat_v1.DeleteRequest{
			Id: id,
		}

		res = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "delete chat service success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, req.Id).Return(nil)
				return mock
			},
		},
		{
			name: "delete chat service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.DeleteChatMock.Expect(ctx, req.Id).Return(serviceErr)
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

			response, err := api.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}
