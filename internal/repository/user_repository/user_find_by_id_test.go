package user_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_FindById тест на получение данных по id
func Test_FindById(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())
	assert.NoError(t, err)

	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	assert.NoError(t, err)

	defer tx.Rollback(ctx)

	prepareData(ctx, t, tx)

	repo := New(tx)

	result, err := repo.FindById(ctx, "00000000-0000-0000-0000-000000000001")

	expected := &entity.User{
		ID:           "00000000-0000-0000-0000-000000000001",
		Email:        "test1@test.ru",
		PasswordHash: "passwordHash1",
	}

	// проверка что репозиторий возвращает корректные данные
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
