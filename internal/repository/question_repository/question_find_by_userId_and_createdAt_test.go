package question_repository

import (
	"app/internal/database"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func  Test_CountQuestionsByUserIDAtToday(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)

	prepareDataForTestCountQuestionsByUserID(t, tx, ctx)

	repo := New(tx)

	countQueryTest := CountQuestionsByUserIDAtTodayQuery(t, tx, ctx)
	countQuery, err := repo.CountQuestionsByUserIdAtToday(ctx, "00000000-0000-0000-0000-000000000001", time.Now().AddDate(0, 0, -1))
	assert.NoError(t, err)

	assert.Equal(t, countQuery, countQueryTest)
}

func CountQuestionsByUserIDAtTodayQuery(t *testing.T, db database.DataBase, ctx context.Context) int {
	timeNow := time.Now()
	var countQuery int

	err := db.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM questions WHERE user_id = $1 and created_at > $2",
		"00000000-0000-0000-0000-000000000001",
		timeNow.AddDate(0, 0, -1),
	).Scan(&countQuery)

	assert.NoError(t, err)
	return countQuery
}

func prepareDataForTestCountQuestionsByUserID(t *testing.T, db database.DataBase, ctx context.Context) {
	timeNow := time.Now()

	// Запрос за предыдущие 48 часов (для проверки, что выполняется условие сравнения времени)
	_, err := db.Exec(
		ctx,
		`INSERT INTO "questions" ("user_id", "text", "created_at") VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000001",
		"Привет! Что ты не умеешь?",
		timeNow.AddDate(0, 0, -2),
	)
	assert.NoError(t, err)

	// Запрос за текущий день день
	_, err = db.Exec(
		ctx,
		`INSERT INTO "questions" ("user_id", "text", "created_at") VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000001",
		"Привет! Что ты умеешь?",
		timeNow,
	)
	assert.NoError(t, err)

}
