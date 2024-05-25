package register_service

import (
	"app/internal/common"
	"app/internal/entity"
	"context"
	"fmt"
)

//go:generate mockgen -destination=mocks/service.go -package=mocks -source=service.go
type Repository interface {
	Create(ctx context.Context, email string, passwordHash string) (*entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, email, password string) (*entity.User, error) {
	passwordHash, err := common.HashPassword(password)

	if err != nil {
		return nil, fmt.Errorf("hash password error %w", err)
	}

	user, err := s.repo.Create(ctx, email, passwordHash)

	if err != nil {
		return nil, fmt.Errorf("could not find user by email %w", err)
	}

	return user, nil
}
