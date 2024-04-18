package premium_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_GetByUserID тест на получение данных о премиум доступе
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

	userTest := getPremium(ctx, t, tx, userId)

	assert.Greater(t, user.ActiveTime, currentTime)
	assert.Equal(t, user.UserID, userId)

	assert.Greater(t, userTest.ActiveTime, currentTime)
	assert.Equal(t, userTest.UserID, userId)

}

// getPremium проверяет наличие пользователя в таблице "Premium" по UUID
func getPremium(ctx context.Context, t *testing.T, db database.DataBase, userID string) *entity.Premium {

	var activeTime time.Time

	err := db.QueryRow(
		ctx,
		`SELECT "active_time" FROM "premium" WHERE "user_id" = $1`,
		userID,
	).Scan(&activeTime)

	assert.NoError(t, err)

	result := entity.Premium{
		UserID:     userID,
		ActiveTime: activeTime,
	}

	return &result
}

// prepareDataForPremium добавляет данные о премиум пользователе в таблицу "premium"
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
