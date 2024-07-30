package login_service

import (
	"app/internal/common"
	"app/internal/entity"
	"app/internal/service/login_service/mocks"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Login(t *testing.T) {
	ctx := context.Background()

	os.Setenv("AUTH_SECRET_KEY", "testsecretkey")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	// создаем мок на репозиторий пользователей
	repoMock := mocks.NewMockRepository(ctrl)

	// описываем поведение мока на момент теста

	password, err := common.HashPassword("passwordHash1")
	assert.NoError(t, err)

	repoMock.
		EXPECT().
		FindByEmail(gomock.Any(), gomock.Any()).
		Return(&entity.User{ // описываем что должен вернуть метод во время теста
			ID:           "00000000-0000-0000-0000-000000000001",
			Email:        "test1@test.ru",
			PasswordHash: password,
		}, nil)

	service := New(repoMock)

	login := LoginDto{Email: "test1@test.ru", Password: "passwordHash1"}

	user, err := service.Login(ctx, login)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.AccessToken)

}
