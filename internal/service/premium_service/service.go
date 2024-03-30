package premium_service

import (
	"app/internal/repository/premium_repository"
	"fmt"
	"time"
)

type Service struct {
	repo *premium_repository.Repository
}

func New(repo *premium_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CheckPremium(userId string) (bool, error) {

	user, err := s.repo.GetByUserID(userId)

	if err != nil {
		return false, fmt.Errorf("question_service checkPremium error:%w", err)
	}

	tmpTime := time.Now()

	if user.ActiveTime.Sub(tmpTime) > 0 {
		return true, nil
	}

	return false, nil
}
