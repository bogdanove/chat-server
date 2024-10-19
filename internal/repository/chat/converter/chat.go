package converter

import (
	serv "github.com/bogdanove/chat-server/internal/model"
	repo "github.com/bogdanove/chat-server/internal/repository/chat/model"
)

// FromServiceToChatRepo конвертация из сервисного слоя в репо
func FromServiceToChatRepo(req *serv.Chat) *repo.Chat {
	return &repo.Chat{
		ChatTitle: req.ChatTitle,
		IDs:       req.IDs,
	}
}

// FromServiceToLogRepo конвертация из сервисного слоя в модель лог репо
func FromServiceToLogRepo(id int64, action string) *repo.ChatLog {
	return &repo.ChatLog{
		ChatID: id,
		Action: action,
	}
}
