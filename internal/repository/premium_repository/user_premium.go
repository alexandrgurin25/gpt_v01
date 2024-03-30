package premium_repository

import (
	"app/internal/entity"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
)

func (r *Repository) GetByUserID(userId string) (*entity.Premium, error) {

	var timeActive time.Time

	err := r.db.QueryRow(
		context.Background(),
		`SELECT "active_time" FROM "premium" WHERE "user_id" = $1`,
		userId,
	).Scan(&timeActive)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("GetByUserID could not get db: %w", err)
	}

	result := &entity.Premium{
		UserID:     userId,
		ActiveTime: timeActive,
	}

	return result, nil
}
