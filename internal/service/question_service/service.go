package question_service

import (
	"app/internal/common"
	"app/internal/entity"
	"log"
	"time"

	"app/internal/repository/question_repository"
)

type AnswerClient interface {
	Request(text string) ([]string, error)
}

type Service struct {
	repo         *question_repository.Repository
	answerClient AnswerClient
}

func New(repo *question_repository.Repository, answerClient AnswerClient) *Service {
	return &Service{repo: repo, answerClient: answerClient}
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

func (s *Service) checkLimit(userId string) error {
	countQuestions, err := s.repo.CountQuestionsByUserIdAtToday(userId, time.Now().AddDate(0, 0, -1))

	if err != nil {
		return err
	}
	
	if countQuestions >= common.MaxQuestionCount {
		log.Printf("У пользователя %s превышен порог запросов: %d > %d", userId, countQuestions, common.MaxQuestionCount)
		return common.InternalError
	}

	return nil
}
