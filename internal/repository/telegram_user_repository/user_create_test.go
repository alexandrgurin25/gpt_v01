package telegram_user_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Create проверка функции question_create
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
	questionCreate, err := repo.CreateUserId(ctx, 25022004)
	assert.NoError(t, err)

	assert.Equal(t, questionCreateTest.ChatId, questionCreate.ChatId)
}

func prepareDataForTestCreate(t *testing.T, ctx context.Context, db database.DataBase) {
	rows, err := db.Query(
		ctx,
		`INSERT INTO "telegram_users" ("user_id", "chat_id") VALUES ($1, $2)`,
		"00000000-0000-0000-0000-000000000001",
		25022004,
	)
	assert.NoError(t, err)

	rows.Close()
}

func getQuestions(t *testing.T, ctx context.Context, db database.DataBase) *entity.TelegramUser {
	var userID string
	var chatID int64

	err := db.QueryRow(
		ctx,
		`SELECT * FROM "telegram_users"`,
	).Scan(&userID, &chatID)

	assert.NoError(t, err)

	question := entity.TelegramUser{
		UserId: userID,
		ChatId: chatID,
	}

	return &question
}
