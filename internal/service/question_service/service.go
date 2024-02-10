package question_service

import (
	"app/internal/common"
	"app/internal/entity"
	"log"
	"time"

	"app/internal/repository/question_repository"
)

type Service struct {
	repo *question_repository.Repository
}

func New(repo *question_repository.Repository) *Service {
	return &Service{repo: repo}
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
	// answer, err := s.answersService.Create(cq)
	return &entity.Answer{}, err
}

func (s *Service) AvailableCount(userId string) (int, error) {
	questions, err := s.repo.FindAll()
	var currentCount int

	if err != nil {
		return 0, err
	}
	// func countQuestions return count questions(int)
	currentCount = common.MaxQuestionCount - s.countAvailableQuestions(questions, userId)

	if currentCount < 0 {
		return 0, nil
	}
	return currentCount, nil
}

func (s *Service) checkLimit(userId string) error {
	questions, err := s.repo.FindAll()

	if err != nil {
		return err
	}

	count := s.countAvailableQuestions(questions, userId)

	if count >= common.MaxQuestionCount {
		log.Printf("У пользователя %s превышен порог запросов: %d > %d", userId, count, common.MaxQuestionCount)
		return common.InternalError
	}

	return nil
}

func (s *Service) countAvailableQuestions(questions []entity.Question, userId string) int {
	userQuestions := s.filterUserId(questions, userId)
	userIntervalQuestions := s.filterTime(userQuestions)

	return len(userIntervalQuestions)
}

func (s *Service) filterUserId(questions []entity.Question, userId string) []entity.Question {
	userQuestions := make([]entity.Question, 0, 0)

	for i := 0; i < len(questions); i++ {
		if userId == questions[i].UserId {
			userQuestions = append(userQuestions, questions[i])
		}
	}

	return userQuestions
}

func (s *Service) filterTime(userQuestions []entity.Question) []entity.Question {

	userIntervalQuestions := make([]entity.Question, 0, 0)

	for i := 0; i < len(userQuestions); i++ {

		createdAtTime, err := time.Parse(common.SQLTimestampFormatTemplate, string((userQuestions[i].CreatedAt)))

		if err != nil {
			log.Printf("%v", err)
			continue
		}

		if time.Now().Unix()-createdAtTime.Unix() < common.QuestionsRateLimitInterval {

			userIntervalQuestions = append(userIntervalQuestions, userQuestions[i])

		}
	}
	return userIntervalQuestions
}
