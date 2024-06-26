package user_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
)

// Update обновляет информацию о пользователе.
func (r *Repository) Update(ctx context.Context, id string, email string, passwordHash string) (*entity.User, error) {
	uuid, err := common.UUIDFromString(id)

	if err != nil {
		return nil, fmt.Errorf("incorrect user id format %w", err)
	}

	_, err = r.db.Exec(
		ctx,
		`UPDATE "users" SET "email" = $2, "password_hash" = $3 WHERE "id" = $1`,
		uuid,
		email,
		passwordHash,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user update error %w", err)
	}

	result := entity.User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
	}

	return &result, nil
}
