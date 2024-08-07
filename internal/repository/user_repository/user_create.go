package user_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

// Создает нового пользователя в таблице "users"
func (r *Repository) Create(ctx context.Context, email string, passwordHash string) (*entity.User, error) {
	var id pgtype.UUID

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO "users" (email, password_hash) VALUES ($1, $2) RETURNING "id"`,
		email,
		passwordHash,
	).Scan(&id)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user create error %w", err)
	}

	result := &entity.User{
		ID:           common.StringFromUUID(&id),
		Email:        email,
		PasswordHash: passwordHash,
	}

	return result, nil
}
