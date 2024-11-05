package tests

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bogdanove/chat-server/internal/api/chat"
	"github.com/bogdanove/chat-server/internal/service"
	serviceMocks "github.com/bogdanove/chat-server/internal/service/mocks"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *chat_v1.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from      = "some sender"
		text      = "some text"
		timeStamp = timestamppb.New(time.Time{})

		mess = &chat_v1.Message{
			From:      from,
			Text:      text,
			Timestamp: timeStamp,
		}

		req = &chat_v1.SendMessageRequest{
			Message: mess,
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
			name: "send message success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
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

			response, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, response)
		})
	}
}
