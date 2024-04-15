package user_repository

import (
	"app/internal/database"
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func Test_Delete(t *testing.T) {
	ctx := context.Background()

	db, err := database.New(database.WithTestConn())

	assert.NoError(t, err)
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)

	assert.NoError(t, err)
	defer tx.Rollback(ctx)

	repo := New(tx)

	id := insertData(ctx, t, tx)

	err = repo.Delete(ctx, id)
	assert.NoError(t, err)

	err = isEmptyDB(ctx, tx)
	assert.NoError(t, err)
}

func isEmptyDB(ctx context.Context, db database.DataBase) error {
	rows, err := db.Query(
		ctx,
		`SELECT * FROM "users"`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return fmt.Errorf("Ожидалось отсутствие строк в таблице, но была найдена хотя бы одна")
	}

	return nil
}

func insertData(ctx context.Context, t *testing.T, db database.DataBase) *pgtype.UUID {

	var id pgtype.UUID

	email := "test@email.ru"

	password := "test123"

	err := db.QueryRow(
		ctx,
		`INSERT INTO "users" (email, password_hash) VALUES ($1, $2)  RETURNING "id"`,
		email,
		password,
	).Scan(&id)

	assert.NoError(t, err)

	return &id
}
