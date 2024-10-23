package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

// DeleteChat - удаление чата из системы по его идентификатору
func (s *Server) DeleteChat(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.DeleteChat(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("chat was deleted for id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
