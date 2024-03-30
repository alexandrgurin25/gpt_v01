package premium_repository

import (
	"app/internal/entity"
	"context"
	"time"
)

func (r *Repository) GetByUserID(userId string) (*entity.Premium, error) {

	var timeActive time.Time

	r.db.QueryRow(
		context.Background(),
		`SELECT "active_time" FROM "premium" WHERE "user_id" = $1 and "active_time" > $2`,
		userId,
		timeActive,
	).Scan(&timeActive)

	result := &entity.Premium{
		UserID:     userId,
		ActiveTime: timeActive,
	}

	return result, nil
}
