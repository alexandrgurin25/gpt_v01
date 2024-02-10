package user_repository

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

// FindAll возвращает всех пользователей.
func (r *Repository) FindAll() ([]entity.User, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT * FROM "users"`,
	)
	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, fmt.Errorf("repository user find all error %w", err)
	}

	users := make([]entity.User, 0, 0)

	for rows.Next() {
		user := entity.User{}
		var id pgtype.UUID

		err = rows.Scan(
			&id,
			&user.Email,
			&user.PasswordHash,
		)

		user.ID = common.StringFromUUID(&id)

		if err != nil {
			log.Printf("%v", err)
			return nil, fmt.Errorf("repository find all delete error %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}
