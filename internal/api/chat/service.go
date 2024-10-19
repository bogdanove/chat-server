package chat

import (
	"github.com/bogdanove/chat-server/internal/service"
	"github.com/bogdanove/chat-server/pkg/chat_v1"
)

// Server - сервер GRPC
type Server struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewServerImplementation - имплементация сервера GRPC
func NewServerImplementation(chatService service.ChatService) *Server {
	return &Server{
		chatService: chatService,
	}
}
