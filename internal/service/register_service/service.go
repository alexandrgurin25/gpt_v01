package register_service

import (
	"app/internal/common"
	"app/internal/entity"
	"app/internal/repository/user_repository"
	"context"
	"fmt"
)

type Service struct {
	repo *user_repository.Repository
}

func New(repo *user_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, registerDto RegisterDto) (*entity.User, error) {
	passwordHash, err := common.HashPassword(registerDto.Password)

	if err != nil {
		return nil, fmt.Errorf("hash password error %w", err)
	}

	user, err := s.repo.Create(ctx, registerDto.Email, passwordHash)

	if err != nil {
		return nil, fmt.Errorf("could not find user by email %w", err)
	}

	return user, nil
}
