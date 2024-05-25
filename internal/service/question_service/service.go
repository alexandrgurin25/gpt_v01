package question_service

import (
	"app/internal/entity"
	"context"
	"time"
)

//go:generate mockgen -destination=mocks/service.go -package=mocks -source=service.go

type AnswerClient interface {
	Request(text string) ([]string, error)
}

type Repository interface {
	Create(ctx context.Context, userId string, text string) (*entity.Question, error)
	CountQuestionsByUserIdAtToday(ctx context.Context, userId string, createdAt time.Time) (int, error)
}

type PremiumService interface {
	CheckPremium(ctx context.Context, userId string) (bool, error)
}

type Service struct {
	repo           Repository
	answerClient   AnswerClient
	premiumService PremiumService

	maxQuestionCount        int
	maxQuestionPremiumCount int
}

func New(
	repo Repository,
	answerClient AnswerClient,
	premiumService PremiumService,
	maxQuestionCount int,
	maxQuestionPremiumCount int,
) *Service {
	return &Service{
		repo:                    repo,
		answerClient:            answerClient,
		premiumService:          premiumService,
		maxQuestionCount:        maxQuestionCount,
		maxQuestionPremiumCount: maxQuestionPremiumCount,
	}
}
