package register_service

import (
	"app/internal/entity"
	"app/internal/service/register_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Test_Register проверяет создание нового пользователя
func Test_Register(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)

	// создаем мок на репозиторий пользователей
	repoMock := mocks.NewMockRepository(ctrl)

	// описываем поведение мока на момент теста
	repoMock.
		EXPECT().
		Create(gomock.Any(), gomock.Any(), gomock.Any()). // вызов метода Create с любыми аргументами
		Return(&entity.User{                              // описываем что должен вернуть метод во время теста
			ID:           "00000000-0000-0000-0000-000000000001",
			Email:        "test1@test.ru",
			PasswordHash: "passwordHash1",
		}, nil)

	// создаем экземпляр сервиса передавая ему в качестве зависимостей мок репозитория
	service := New(repoMock)

	// вызов проверяемого метода
	user, err := service.Register(ctx, "test1@test.ru", "password")

	// проверка возвращаемых значений
	assert.NoError(t, err)
	assert.Equal(t, "00000000-0000-0000-0000-000000000001", user.ID)
	assert.Equal(t, "test1@test.ru", user.Email)
	assert.Equal(t, "passwordHash1", user.PasswordHash)
}
