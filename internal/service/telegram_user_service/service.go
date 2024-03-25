package telegram_user_service

import (
	"app/internal/entity"
	"app/internal/repository/telegram_user_repository"
	"fmt"
)

type Service struct {
	repo *telegram_user_repository.Repository
}

func New(repo *telegram_user_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUserIdByChatId(chatId int64) (*entity.TelegramUser, error) {
	var err error
	
	user, err := s.repo.GetUserId(chatId)

	if user.UserId == ""{
		user, err = s.repo.CreateUserId(chatId)
	}

	if err != nil {
		return nil, fmt.Errorf("could not add telegram user in db %w", err)
	}

	return user, nil
}
