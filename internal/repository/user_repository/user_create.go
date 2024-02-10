package user_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

// Create создает нового пользователя.
func (r *Repository) Create(email string, passwordHash string) (*entity.User, error) {
	var id pgtype.UUID

	err := r.db.QueryRow(
		context.Background(),
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
