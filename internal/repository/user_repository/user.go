package user_repository

import (
	"app/internal/database"
)

// repository представляет репозиторий пользователей.
type Repository struct {
	db *database.DataBase
}

// New создает новый экземпляр repository.
func New(db *database.DataBase) *Repository {
	return &Repository{db: db}
}
