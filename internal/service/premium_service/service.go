package premium_service

import (
	"app/internal/repository/premium_repository"
	"context"
	"fmt"
	"time"
)

type Service struct {
	repo *premium_repository.Repository
}

func New(repo *premium_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CheckPremium(ctx context.Context, userId string) (bool, error) {

	user, err := s.repo.GetByUserID(ctx, userId)

	if err != nil {
		return false, fmt.Errorf("question_service checkPremium error: %w", err)
	}

	currentTime := time.Now()

	if user != nil && user.ActiveTime.Sub(currentTime) > 0 {
		return true, nil
	}

	return false, nil
}
