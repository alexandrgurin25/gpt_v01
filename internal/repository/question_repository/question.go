package question_repository

import (
	"app/internal/database"
)

type Repository struct {
	db database.DataBase
}

func New(db database.DataBase) *Repository {
	return &Repository{db}
}
