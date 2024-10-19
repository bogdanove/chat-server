package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

// SendMessage - отправка сообщения на сервер
func (s *Server) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("incoming message received: %+v", req.Message)

	return &emptypb.Empty{}, nil
}
