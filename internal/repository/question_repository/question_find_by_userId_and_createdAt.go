package question_repository

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Подсчет количества количества вопросов от пользователя за n-e время
func (r *Repository) CountQuestionsByUserIdAtToday(ctx context.Context, userId string, createdAt time.Time) (int, error) {
	var countQuery int

	err := r.db.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM questions WHERE user_id = $1 and created_at > $2",
		userId,
		createdAt,
	).Scan(&countQuery)

	if err != nil {
		log.Printf("%v", err)
		return 0, fmt.Errorf("repository question FindByUserIdAndCreatedAt error %w", err)
	}

	return countQuery, err
}
