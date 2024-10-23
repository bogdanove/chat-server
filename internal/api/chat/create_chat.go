package chat

import (
	"context"
	"log"

	"github.com/bogdanove/chat-server/internal/converter"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
	"github.com/pkg/errors"
)

// CreateChat - создания нового чата
func (s *Server) CreateChat(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}

	id, err := s.chatService.CreateChat(ctx, converter.FromProtoToService(req))
	if err != nil {
		return nil, err
	}

	log.Printf("new chat was created with id: %d", id)

	return &chat_v1.CreateResponse{
		Id: id,
	}, nil
}
