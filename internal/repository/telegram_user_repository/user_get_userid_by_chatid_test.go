package telegram_user_repository

import (
	"app/internal/database"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserId(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	prepareDataForTelegram(t, ctx, tx)
	repo := New(tx)

	userTelegram, err := repo.GetUserId(ctx, 25022004)
	assert.NoError(t, err)

	assert.Equal(t, int64(25022004), userTelegram.ChatId)
	assert.Equal(t, userTelegram.UserId, "00000000-0000-0000-0000-000000000001")
}

func prepareDataForTelegram(t *testing.T, ctx context.Context, db database.DataBase) {
	_, err := db.Exec(
		ctx,
		`INSERT INTO "telegram_users" (user_id, chat_id) VALUES($1, $2)`,
		"00000000-0000-0000-0000-000000000001",
		25022004,
	)
	assert.NoError(t, err)
}
