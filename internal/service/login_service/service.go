package login_service

import (
	"app/internal/common"
	"app/internal/repository/user_repository"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	repo *user_repository.Repository
}

func New(repo *user_repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(loginDto LoginDto) (*AuthDto, error) {
	// Находим пользователя по электронной почте
	user, err := s.repo.FindByEmail(loginDto.Email)

	if err != nil {
		return nil, fmt.Errorf("incorrect email or password")
	}

	// Проверяем совпадение паролей
	if !common.CheckPasswordHash(loginDto.Password, user.PasswordHash) {
		return nil, fmt.Errorf("incorrect email or password")
	}

	// Создаем AccessToken
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"iat": time.Now().Unix(),
	})

	accessToken, err := token.SignedString(common.SecretKey)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	authDto := &AuthDto{
		AccessToken: accessToken,
	}
	// Возвращаем AuthDto с AccessToken
	return authDto, nil
}
