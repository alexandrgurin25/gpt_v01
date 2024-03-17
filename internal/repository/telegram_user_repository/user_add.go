package telegram_user_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Repository struct {
	db *database.DataBase
}

func New(db *database.DataBase) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrUpdateUserId(chatId int64) (*entity.TelegramUser, error) {

	userId := uuid.New()

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO "telegram_users" (user_id, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO UPDATE SET user_id = EXCLUDED.user_id RETURNING "user_id"`,
		userId,
		chatId,
	).Scan(&userId)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error %w", err)
	}

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatId,
	}

	return result, nil
}
