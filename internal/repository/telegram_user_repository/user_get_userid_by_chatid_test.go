package telegram_user_repository

import (
	"app/internal/common"
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
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

	testUserTelegram := getUserIdForTest(t, ctx, tx, 25022004)

	userTelegram, err := repo.GetUserId(ctx, 25022004)
	assert.NoError(t, err)

	assert.Equal(t, testUserTelegram, userTelegram)
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

func getUserIdForTest(t *testing.T, ctx context.Context, db database.DataBase, chatID int64) *entity.TelegramUser {
	var userID *pgtype.UUID

	err := db.QueryRow(
		ctx,
		`SELECT user_id FROM telegram_users WHERE chat_id = $1`,
		chatID,
	).Scan(&userID)

	assert.NoError(t, err)

	result := entity.TelegramUser{
		UserId: common.StringFromUUID(userID),
		ChatId: chatID,
	}

	return &result
}
