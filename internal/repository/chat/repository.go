package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogdanove/chat-server/internal/client/db"
	"github.com/bogdanove/chat-server/internal/repository"
	"github.com/bogdanove/chat-server/internal/repository/chat/model"
)

const (
	chatTableName      = "chat"
	chatUsersTableName = "chat_users"
	chatLogTableName   = "chat_log"

	chatIDColumn    = "id"
	chatTitleColumn = "chat_title"

	chatUsersChatIDColumn = "chat_id"
	chatUsersUserIDColumn = "user_id"

	chatLogChatIDColumn = "chat_id"
	chatLogActionColumn = "action"
)

type chatRepo struct {
	db db.Client
}

// NewChatRepository - создание репозитория для чата
func NewChatRepository(db db.Client) repository.ChatRepository {
	return &chatRepo{db: db}
}

// CreateChat - создания нового чата
func (r *chatRepo) CreateChat(ctx context.Context, req *model.Chat) (int64, error) {
	log.Printf("create new chat with title: %v", req.ChatTitle)

	queryChat, args, err := sq.Insert(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatTitleColumn).
		Values(req.ChatTitle).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}

	qc := db.Query{
		Name:     "chat_repository.Create_chat",
		QueryRaw: queryChat,
	}

	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, qc, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to insert chat: %v", err)
		return 0, err
	}

	builderInsertChatUser := sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatUsersChatIDColumn, chatUsersUserIDColumn)
	for _, id := range req.IDs {
		builderInsertChatUser = builderInsertChatUser.Values(chatID, id)
	}

	queryChatUser, argsUsr, err := builderInsertChatUser.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}

	qu := db.Query{
		Name:     "chat_repository.Create_users",
		QueryRaw: queryChatUser,
	}

	_, err = r.db.DB().ExecContext(ctx, qu, argsUsr...)
	if err != nil {
		log.Printf("failed to insert chat_users: %v", err)
		return 0, err
	}

	return chatID, nil
}

// DeleteChat - удаление чата из системы по его идентификатору
func (r *chatRepo) DeleteChat(ctx context.Context, req int64) error {

	queryDelChatUser, argsUsr, err := sq.Delete(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatUsersChatIDColumn: req}).ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	qu := db.Query{
		Name:     "chat_repository.Delete_users",
		QueryRaw: queryDelChatUser,
	}

	_, err = r.db.DB().ExecContext(ctx, qu, argsUsr...)
	if err != nil {
		log.Printf("failed to delete chat_user: %v", err)
		return err
	}

	queryDelChat, args, err := sq.Delete(chatTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatIDColumn: req}).ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	qc := db.Query{
		Name:     "chat_repository.Delete_chat",
		QueryRaw: queryDelChat,
	}

	_, err = r.db.DB().ExecContext(ctx, qc, args...)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return err
	}

	return nil
}

// SaveLog - сохранение записи о действиях пользователя
func (r *chatRepo) SaveLog(ctx context.Context, req *model.ChatLog) error {
	log.Printf("create new chat_log with chat_id: %d", req.ChatID)

	queryChatLog, args, err := sq.Insert(chatLogTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatLogChatIDColumn, chatLogActionColumn).
		Values(req.ChatID, req.Action).ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	ql := db.Query{
		Name:     "chat_repository.SaveLog",
		QueryRaw: queryChatLog,
	}

	_, err = r.db.DB().ExecContext(ctx, ql, args...)
	if err != nil {
		log.Printf("failed to insert chat_log: %v", err)
		return err
	}
	return nil
}
