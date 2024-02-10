package question_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) Create(userId string, text string) (*entity.Question, error) {
	var id int64
	var createdAt time.Time

	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO "questions" ("user_id", "text") VALUES ($1, $2) RETURNING "id", "created_at"`,
		userId,
		text,
	).Scan(&id, &createdAt)

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
