package premium_repository

import (
	"app/internal/database"
	"context"
	"time"
)

type Repository struct {
	db *database.DataBase
}

func New(db *database.DataBase) *Repository {
	return &Repository{db}
}

func (r *Repository) HasByUserIDAndActiveTime(userId string) (time.Time, error) {

	var timeActive time.Time

	r.db.QueryRow(
		context.Background(),
		`SELECT active_time FROM "premium" WHERE "user_id" = $1 and "active_time" > $2`,
		userId,
		timeActive,
	).Scan(&timeActive)

	return timeActive, nil
}
