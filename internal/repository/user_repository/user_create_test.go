package user_repository

import (
	"app/internal/common"
	"app/internal/database"
	"app/internal/entity"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

// тест на получение всех данных
func Test_Create(t *testing.T) {
	ctx := context.Background()

	// создаем соединение к бд для теста
	db, err := database.New(database.WithTestConn())
	// проверяем, что соединение было создано без ошибки
	assert.NoError(t, err)
	defer db.Close(ctx)

	// каждый тест запускаем отдельной транзакцией в БД
	tx, err := db.Begin(ctx)
	assert.NoError(t, err)

	// после теста транзакцию откатываем, чтобы в Бд ничего не сохранилось
	defer tx.Rollback(ctx)

	// инициализация репозитория
	repo := New(tx)

	// вызов тестируемого метода
	result, err := repo.Create(ctx, "test@test.ru", "password_hash")
	assert.NoError(t, err)

	// проверка что репозиторий возвращает корректные данные
	assert.Greater(t, len(result.ID), 0)
	assert.Equal(t, "test@test.ru", result.Email)
	assert.Equal(t, "password_hash", result.PasswordHash)

	// получение данных из бд
	dataInDB := getData(ctx, t, tx)

	// проверка что в базу вставлены корректные данные
	assert.Equal(t, 1, len(dataInDB))
	assert.Equal(t, result.ID, dataInDB[0].ID)
	assert.Equal(t, "test@test.ru", dataInDB[0].Email)
	assert.Equal(t, "password_hash", dataInDB[0].PasswordHash)
}

// функция для получения вставленных данных
func getData(ctx context.Context, t *testing.T, db database.DataBase) []entity.User {
	var users []entity.User

	rows, err := db.Query(
		ctx,
		`SELECT "id", "email", "password_hash" FROM "users"`,
	)
	assert.NoError(t, err)

	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		var id pgtype.UUID

		err = rows.Scan(
			&id,
			&user.Email,
			&user.PasswordHash,
		)

		assert.NoError(t, err)

		user.ID = common.StringFromUUID(&id)

		users = append(users, user)
	}

	return users
}
