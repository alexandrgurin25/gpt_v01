package premium_repository

import (
	"app/internal/database"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// тест на получение данных о премиум доступе
func Test_GetByUserID(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	// Тестовый UUID пользователя
	userId := "00000000-0000-0000-0000-000000000001"
	currentTime := time.Now()

	prepareDataForPremium(ctx, t, tx, userId)

	repo := New(tx)
	assert.NoError(t, err)

	user, err := repo.GetByUserID(ctx, userId)
	assert.NoError(t, err)

	assert.Greater(t, user.ActiveTime, currentTime)
	assert.Equal(t, user.UserID, userId)
}

func prepareDataForPremium(ctx context.Context, t *testing.T, db database.DataBase, userID string) {
	// Добавляем к текущей дате +1 день
	accessPeriod := time.Now().AddDate(0, 0, 1)

	_, err := db.Exec(
		ctx,
		`INSERT INTO "premium" (user_id, active_time) VALUES ($1, $2)`,
		userID,
		accessPeriod,
	)

	assert.NoError(t, err)
}
