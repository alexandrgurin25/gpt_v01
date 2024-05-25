package question_service

import (
	"context"
	"time"
)

func (s *Service) AvailableCount(ctx context.Context, userId string) (int, error) {
	countQuestions, err := s.repo.CountQuestionsByUserIdAtToday(ctx, userId, time.Now().AddDate(0, 0, -1))
	var currentCount int

	if err != nil {
		return 0, err
	}

	hasPremium, err := s.premiumService.CheckPremium(ctx, userId)

	if err != nil {
		return 0, err
	}

	questionsLimit := s.maxQuestionCount

	if hasPremium {
		questionsLimit = s.maxQuestionPremiumCount
	}

	currentCount = questionsLimit - countQuestions

	if currentCount < 0 {
		return 0, nil
	}
	return currentCount, nil
}
