package user_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"
)

// FindById возвращает пользователя по идентификатору.
func (r *Repository) FindById(id string) (*entity.User, error) {
	uuid, err := common.UUIDFromString(id)

	if err != nil {
		return nil, fmt.Errorf("incorrect user id format %w", err)
	}

	user := entity.User{}

	err = r.db.QueryRow(
		context.Background(),
		`SELECT * FROM "users" WHERE "id" = $1`,
		uuid,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user find by id error %w", err)
	}

	return &user, nil
}
