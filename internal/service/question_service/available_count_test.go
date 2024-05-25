package question_service

import (
	"app/internal/service/question_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Test_AvailableCount - тест на метод подсчета доступных для пользователя вопросов
func Test_AvailableCount(t *testing.T) {
	// описываем структуры с тейсткейсами
	tests := []struct {
		name                          string
		countQuestionsByUserIdAtToday int
		hasPremium                    bool
		answerClientResult            []string
		expected                      int
	}{
		{
			name:                          "Пользователь не задавал вопросы и у него нет премиума",
			countQuestionsByUserIdAtToday: 0,
			hasPremium:                    false,
			expected:                      10,
		},
		{
			name:                          "Пользователь уже задал 3 вопроса и у него нет премиума",
			countQuestionsByUserIdAtToday: 3,
			hasPremium:                    false,
			expected:                      7,
		},
		{
			name:                          "Пользователь не задавал вопросы и у него есть премиум",
			countQuestionsByUserIdAtToday: 0,
			hasPremium:                    true,
			expected:                      20,
		},
		{
			name:                          "Пользователь уже задал 3 вопроса и у него нет премиума",
			countQuestionsByUserIdAtToday: 3,
			hasPremium:                    true,
			expected:                      17,
		},
	}

	for _, test := range tests {
		// запускаем каждый тест кейс
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)

			// создаем мок на репозиторий для вопросов
			repoMock := mocks.NewMockRepository(ctrl)

			// определяем, что должен вернуть метод репозитория и сколько раз быть вызван
			repoMock.
				EXPECT().                                                                // ожидаемое поведение
				CountQuestionsByUserIdAtToday(gomock.Any(), gomock.Any(), gomock.Any()). // gomock.Any() - любое значение
				Times(1).                                                                // будет вызван 1 раз
				Return(test.countQuestionsByUserIdAtToday, nil)                          // вернет значение из test.countQuestionsByUserIdAtToday и ошибку nil

			// создаем мок на сервис премиума
			premiumServiceMock := mocks.NewMockPremiumService(ctrl)

			// определяем, что должен вернуть сервис премиума и сколько раз быть вызван
			premiumServiceMock.
				EXPECT().
				CheckPremium(gomock.Any(), gomock.Any()).
				Times(1).
				Return(test.hasPremium, nil)

			// создаем мок на клиент к gpt
			answerClientMock := mocks.NewMockAnswerClient(ctrl)

			// в AvailableCount клиент не вызывается, поэтому в Times указываем - 0
			answerClientMock.
				EXPECT().
				Request(gomock.Any()).
				Times(0) // будет проверено, что этот метод не вызывался

			// передаем в сервис в качестве зависимостей моки
			service := New(repoMock, answerClientMock, premiumServiceMock, 10, 20)

			// вызываем тестируемый метод
			result, err := service.AvailableCount(ctx, "00000000-0000-0000-0000-000000000001")

			// проверяем возвращаемые результаты
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
