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

func (r *Repository) CreateUserId(chatID int64) (*entity.TelegramUser, error) {

	userId := uuid.New()

	_, err := r.db.Query(
		context.Background(),
		`INSERT INTO "telegram_users" (user_id, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO NOTHING`,
		userId,
		chatID,
	)

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

func (r *Repository) GetUserId(chatID int64) (*entity.TelegramUser, error) {

	var userId uuid.UUID
	var hasChatID bool

	err := r.db.QueryRow(
		context.Background(),
		`SELECT EXISTS(SELECT * FROM "telegram_users" WHERE "chat_id" = $1)`,
		chatID,
	).Scan(&hasChatID)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("telegram_repository GetUserID error: %w", err)
	}

	if hasChatID {
		err = r.db.QueryRow(
			context.Background(),
			`SELECT user_id FROM "telegram_users" WHERE "chat_id" = $1`,
			chatID,
		).Scan(&userId)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository user create error %w", err)
		}

	}

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatID,
	}

	return result, nil
}
