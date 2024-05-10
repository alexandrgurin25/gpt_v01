package question_repository

import (
	"app/internal/database"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_Count проверяет Count на корректную работу
func Test_Count(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	repo := New(tx)

	prepareData(t, ctx, db)

	countQuery, err := repo.Count(ctx)
	assert.NoError(t, err)

	assert.Equal(t, countQuery, 3)

}

/*
Подготавливаем таблицу questions так, чтобы в ней были записи от 2 разных клиентов.
Клиент N оставил запрос 2 раза, клиент M оставил запрос 1 раз. Всего 3 запроса
*/		
func prepareData(t *testing.T, ctx context.Context, db database.DataBase) {
	currentTime := time.Now()

	// 1-я Запись клиента N
	_, err := db.Exec(
		ctx,
		`INSERT INTO "questions" (user_id, text, created_at) VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000001",
		"Test",
		currentTime,
	)
	assert.NoError(t, err)

	// 2-я Запись клиента N
	_, err = db.Exec(
		ctx,
		`INSERT INTO "questions" (user_id, text, created_at) VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000001",
		"Test",
		currentTime,
	)
	assert.NoError(t, err)

	// 1=я Запись клиента M
	currentTime = currentTime.AddDate(0, 0, 1)
	_, err = db.Exec(
		ctx,
		`INSERT INTO "questions" (user_id, text, created_at) VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000002",
		"Test",
		currentTime,
	)
	assert.NoError(t, err)
}
