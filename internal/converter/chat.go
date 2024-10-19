package converter

import (
	"github.com/bogdanove/chat-server/internal/model"
	req "github.com/bogdanove/chat-server/pkg/chat_v1"
)

// FromProtoToService - конвертация из протобаф в модель сервисного слоя
func FromProtoToService(req *req.CreateRequest) *model.Chat {
	return &model.Chat{
		ChatTitle: req.ChatTitle,
		IDs:       req.Ids,
	}
}
