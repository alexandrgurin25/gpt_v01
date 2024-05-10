package user_repository

import (
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_FindAll - получение пользователей из базы данных
func Test_FindAll(t *testing.T) {
	ctx := context.Background()

	// создаем соединение к бд для теста
	db, err := database.New(database.WithTestConn())
	// проверяем, что соединение было создано без ошибки
	assert.NoError(t, err)
	defer db.Close(ctx) // добавляем отложенные закрытие соединения после теста

	// каждый тест запускаем отдельной транзакцией в БД
	tx, err := db.Begin(ctx)
	assert.NoError(t, err)

	// после теста транзакцию откатываем, чтобы в Бд ничего не сохранилось
	defer tx.Rollback(ctx)

	// подготовка данных перед тестом
	prepareData(ctx, t, tx)

	// инициализация репозитория
	repo := New(tx)

	// вызов тестируемого метода
	result, err := repo.FindAll(ctx)

	// описание ожидаемого результата
	expected := []entity.User{
		{
			ID:           "00000000-0000-0000-0000-000000000001",
			Email:        "test1@test.ru",
			PasswordHash: "passwordHash1",
		},
	}

	// проверка результатов
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func prepareData(ctx context.Context, t *testing.T, db database.DataBase) {
	_, err := db.Exec(
		ctx,
		`INSERT INTO "users" (id, email, password_hash) VALUES ($1, $2, $3)`,
		"00000000-0000-0000-0000-000000000001",
		"test1@test.ru",
		"passwordHash1",
	)

	assert.NoError(t, err)
}
