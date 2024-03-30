package telegram_user_repository

import (
	"app/internal/entity"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

func (r *Repository) GetUserId(ctx context.Context, chatID int64) (*entity.TelegramUser, error) {

	var userId uuid.UUID

	err := r.db.QueryRow(
		ctx,
<<<<<<< HEAD
		`SELECT user_id FROM "telegram_users" WHERE "chat_id" = $1`,
=======
		`SELECT (SELECT user_id FROM "telegram_users" WHERE "chat_id" = $1) AS id`,
>>>>>>> c1bee4d (added integration tests)
		chatID,
	).Scan(&userId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("repository GetUserId error: %w", err)
	}

	result := &entity.TelegramUser{
		UserId: userId.String(),
		ChatId: chatID,
	}

	return result, nil
}
