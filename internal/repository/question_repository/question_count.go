package question_repository

import (
	"context"
	"fmt"
	"log"
)

// Count подсчет количества всех вопросов от пользователей
func (r *Repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM "questions"`,
	).Scan(&count)

	if err != nil {
		log.Printf("%v", err)
		return 0, fmt.Errorf("repository question count error %w", err)
	}
	return count, nil
}
