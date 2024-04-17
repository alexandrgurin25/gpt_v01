package telegram_user_repository

import (
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Добавление пользователя из telegram в таблицу "telegram_users"
func (r *Repository) CreateUserId(ctx context.Context, chatID int64) (*entity.TelegramUser, error) {

	userId := uuid.New()

	rows, err := r.db.Query(
		ctx,
		`INSERT INTO "telegram_users" (user_id, chat_id) VALUES ($1, $2)`,
		userId,
		chatID,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error: %w", err)
	}

	defer rows.Close()

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatID,
	}

	return result, nil
}
