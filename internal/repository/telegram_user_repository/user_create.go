package telegram_user_repository

import (
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (r *Repository) CreateUserId(chatID int64) (*entity.TelegramUser, error) {

	userId := uuid.New()

	rows, err := r.db.Query(
		context.Background(),
		`INSERT INTO "telegram_users" (user_id, chat_id) VALUES ($1, $2)`,
		userId,
		chatID,
	)

	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error: %w", err)
	}

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatID,
	}

	return result, nil
}
