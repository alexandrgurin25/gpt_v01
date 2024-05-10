package user_repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// Delete удаляет пользователя по идентификатору.
func (r *Repository) Delete(ctx context.Context, id *pgtype.UUID) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM "users" WHERE "id" = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("repository user delete error %w", err)
	}

	return nil
}
