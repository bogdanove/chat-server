package converter

import (
	"errors"

	"github.com/bogdanove/chat-server/internal/model"
	req "github.com/bogdanove/chat-server/pkg/chat_v1"
)

// FromProtoToService - конвертация из протобаф в модель сервисного слоя
func FromProtoToService(req *req.CreateRequest) (*model.Chat, error) {
	if req == nil {
		return nil, errors.New("request for convert is nil")
	}
	return &model.Chat{
		ChatTitle: req.ChatTitle,
		IDs:       req.Ids,
	}, nil
}
