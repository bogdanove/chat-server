package model

// Chat - структура чат для репо слоя
type Chat struct {
	IDs       []int64
	ChatTitle string
}

// ChatLog - структура чат лога для сохранения действий пользователя
type ChatLog struct {
	ChatID int64
	Action string
}
