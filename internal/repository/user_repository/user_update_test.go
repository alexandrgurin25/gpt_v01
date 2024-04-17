package user_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Update(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	repo := New(tx)

	prepareData(ctx, t, tx)

	userID := "00000000-0000-0000-0000-000000000001"

	result, err := repo.Update(ctx, userID, "test123@email.ru", "pass123")
	assert.NoError(t, err)

	// проверка что репозиторий возвращает корректные данные
	assert.Greater(t, len(result.ID), 0)
	assert.Equal(t, "test123@email.ru", result.Email)
	assert.Equal(t, "pass123", result.PasswordHash)

	// Создать отдельную "getData" by id
	dataInDB := getDataById(ctx, userID, t, tx)

	// проверка что в базу вставлены корректные данные
	assert.Equal(t, result.ID, dataInDB.ID)
	assert.Equal(t, "test123@email.ru", dataInDB.Email)
	assert.Equal(t, "pass123", dataInDB.PasswordHash)
}

func getDataById(ctx context.Context, userID string, t *testing.T, db database.DataBase) *entity.User {
	var user entity.User

	err := db.QueryRow(
		ctx,
		`SELECT "id", "email", "password_hash" FROM "users" WHERE "id" = $1`,
		userID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash)

	assert.NoError(t, err)

	return &user
}
