package question_service

import (
	"app/internal/common"
	"app/internal/entity"
	"fmt"
	"log"
	"time"

	"app/internal/repository/premium_repository"
	"app/internal/repository/question_repository"
)

type AnswerClient interface {
	Request(text string) ([]string, error)
}

type Service struct {
	repo         *question_repository.Repository
	answerClient AnswerClient
	repoPremium  *premium_repository.Repository
}

func New(repo *question_repository.Repository, answerClient AnswerClient, repoPremium *premium_repository.Repository) *Service {
	return &Service{repo: repo, answerClient: answerClient, repoPremium: repoPremium}
}

func (s *Service) Create(userId string, text string) (*entity.Answer, error) {
	err := s.checkLimit(userId)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.Create(userId, text)
	if err != nil {
		return nil, err
	}

	answer, err := s.answerClient.Request(text)
	if err != nil {
		return nil, err
	}
	return &entity.Answer{Texts: answer}, err
}

func (s *Service) AvailableCount(userId string) (int, error) {
	countQuestions, err := s.repo.CountQuestionsByUserIdAtToday(userId, time.Now().AddDate(0, 0, -1))
	var currentCount int

	if err != nil {
		return 0, err
	}

	currentCount = common.MaxQuestionCount - countQuestions

	if currentCount < 0 {
		return 0, nil
	}
	return currentCount, nil
}

func (s *Service) checkPremium(userId string) (bool, error) {

	timePremium, err := s.repoPremium.HasByUserIDAndActiveTime(userId)

	if err != nil {
		return false, fmt.Errorf("question_service checkPremium error:%w", err)
	}

	tmpTime := time.Now()

	if timePremium.Sub(tmpTime) > 0 {
		return true, nil
	}

	return false, nil
}

func (s *Service) checkLimit(userId string) error {
	countQuestions, err := s.repo.CountQuestionsByUserIdAtToday(userId, time.Now().AddDate(0, 0, -1))

	if err != nil {
		return fmt.Errorf("question_service checkLimit error: %w", err)
	}

	checkPremium, err := s.checkPremium(userId)

	if err != nil {
		return err
	}

	if checkPremium {
		if countQuestions >= common.MaxQuestionCountPremium {
			log.Printf("У пользователя %s c ПРЕМИУМ доступом превышен порог запросов: %d > %d", userId, countQuestions, common.MaxQuestionCount)
			return common.InternalError
		}
	} else {
		if countQuestions >= common.MaxQuestionCount {
			log.Printf("У пользователя %s превышен порог запросов: %d > %d", userId, countQuestions, common.MaxQuestionCount)
			return common.InternalError
		}
	}

	return nil
}
