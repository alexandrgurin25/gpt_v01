package question_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) FindAll() ([]entity.Question, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT * FROM "questions"`,
	)
	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository question find all error %w", err)
	}

	questions := make([]entity.Question, 0, 0)

	for rows.Next() {
		var createdAt time.Time

		question := entity.Question{}

		err = rows.Scan(
			&question.ID,
			&question.UserId,
			&question.Text,
			&createdAt,
		)
		// Преобразование в с помощью метода Format()
		timeStr := createdAt.Format(common.SQLTimestampFormatTemplate)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository question find all error %w", err)
		}
		question.CreatedAt = timeStr

		questions = append(questions, question)
	}

	return questions, err
}
