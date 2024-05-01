package question_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
	"time"
)

// Добавление в таблицу "questions" вопроса от пользователя
func (r *Repository) Create(ctx context.Context, userId string, text string) (*entity.Question, error) {
	var createdAt time.Time

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO "questions" ("user_id", "text") VALUES ($1, $2) RETURNING "created_at"`,
		userId,
		text,
	).Scan(&createdAt)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository question create error %w", err)
	}

	// Преобразование в строку с помощью метода Format()
	timeStr := createdAt.Format(common.SQLTimestampFormatTemplate)

	result := entity.Question{
		UserId:    userId,
		Text:      text,
		CreatedAt: timeStr,
	}

	return &result, nil
}
