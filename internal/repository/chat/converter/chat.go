package converter

import (
	repo "github.com/bogdanove/chat-server/internal/repository/chat/model"
)

// FromServiceToLogRepo конвертация из сервисного слоя в модель лог репо
func FromServiceToLogRepo(id int64, action string) *repo.ChatLog {
	return &repo.ChatLog{
		ChatID: id,
		Action: action,
	}
}
