package question_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Test_Create проверка функции question_create
func Test_Create(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	prepareDataForTestCreate(t, ctx, tx)

	repo := New(tx)
	questionCreateTest := getQuestions(t, ctx, tx)
	questionCreate, err := repo.Create(ctx, "00000000-0000-0000-0000-000000000001", "Привет! Что ты умеешь?")
	assert.NoError(t, err)

	assert.Equal(t, questionCreate.UserId, "00000000-0000-0000-0000-000000000001")
	assert.Equal(t, questionCreateTest.UserId,"00000000-0000-0000-0000-000000000001")
	assert.Equal(t, "Привет! Что ты умеешь?", questionCreate.Text)
	assert.Equal(t, "Привет! Что ты умеешь?", questionCreateTest.Text)
	assert.Equal(t, questionCreate.Text, questionCreate.Text)
}

func prepareDataForTestCreate(t *testing.T, ctx context.Context, db database.DataBase) {
	rows, err := db.Query(
		ctx,
		`INSERT INTO "questions" ("user_id", "text") VALUES ($1, $2)`,
		"00000000-0000-0000-0000-000000000001",
		"Привет! Что ты умеешь?",
	)
	assert.NoError(t, err)
	
	rows.Close()
}

func getQuestions(t *testing.T, ctx context.Context, db database.DataBase) *entity.Question {
	var userID, text string
	var createdAt time.Time

	err := db.QueryRow(
		ctx,
		`SELECT "user_id", "text", "created_at" FROM "questions"`,
	).Scan(&userID, &text, &createdAt)

	assert.NoError(t, err)

	question := entity.Question{
		UserId:    userID,
		Text:      text,
		CreatedAt: createdAt.String(),
	}

	return &question
}
