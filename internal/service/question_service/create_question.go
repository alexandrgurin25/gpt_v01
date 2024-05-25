package question_service

import (
	"context"
	"fmt"
	"time"

	"app/internal/entity"
)

func (s *Service) CreateQuestion(ctx context.Context, userId string, text string) (*entity.Answer, error) {
	err := s.checkLimit(ctx, userId)
	if err != nil {
		arrString := make([]string, 2)
		arrString[0] = "У вас превышен порог запросов за последние 24ч"
		arrString[1] = "Обратитесь к @alexan_25"
		return &entity.Answer{Texts: arrString}, nil
	}
	_, err = s.repo.Create(ctx, userId, text)
	if err != nil {
		return nil, err
	}

	answer, err := s.answerClient.Request(text)
	if err != nil {
		return nil, err
	}
	return &entity.Answer{Texts: answer}, err
}

func (s *Service) checkLimit(ctx context.Context, userId string) error {
	countQuestions, err := s.repo.CountQuestionsByUserIdAtToday(ctx, userId, time.Now().AddDate(0, 0, -1))

	if err != nil {
		return fmt.Errorf("question_service checkLimit error: %w", err)
	}

	checkPremium, err := s.premiumService.CheckPremium(ctx, userId)

	if err != nil {
		return err
	}

	if checkPremium {
		if countQuestions >= s.maxQuestionPremiumCount {
			return fmt.Errorf("У пользователя %s c ПРЕМИУМ доступом превышен порог запросов: %d > %d", userId, countQuestions, s.maxQuestionPremiumCount)
		}
	} else {
		if countQuestions >= s.maxQuestionCount {
			return fmt.Errorf("У пользователя %s превышен порог запросов: %d > %d", userId, countQuestions, s.maxQuestionCount)
		}
	}

	return nil
}
