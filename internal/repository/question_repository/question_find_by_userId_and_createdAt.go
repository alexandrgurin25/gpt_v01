package question_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *Repository) CountQuestionsByUserIdAtToday(userId string, createdAt time.Time) (int, error) {
	var countQuery int
	
	err := r.db.QueryRow(
		context.Background(),
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

func (r *Repository) FindAll() ([]entity.Question, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT * FROM "questions"`,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository question find all error %w", err)
	}

	defer rows.Close()

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
