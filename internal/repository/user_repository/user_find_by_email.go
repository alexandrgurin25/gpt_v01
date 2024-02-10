package user_repository

import (
	"app/internal/entity"
	"context"
	"fmt"
	"log"
)

// FindByEmail возвращает пользователя по адресу электронной почты.
func (r *Repository) FindByEmail(email string) (*entity.User, error) {
	user := entity.User{}

	err := r.db.QueryRow(
		context.Background(),
		`SELECT * FROM "users" WHERE "email" = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user find by email error %w", err)
	}

	return &user, nil
}
