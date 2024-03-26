package telegram_user_repository

import (
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (r *Repository) GetUserId(chatID int64) (*entity.TelegramUser, error) {

	var userId uuid.UUID

	err := r.db.QueryRow(
		context.Background(),
		`SELECT (SELECT user_id FROM "telegram_users" WHERE "chat_id" = $1) AS id`,
		chatID,
	).Scan(&userId)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error %w", err)
	}

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatID,
	}

	return result, nil
}
