package question_service

import (
	"app/internal/entity"
	"app/internal/service/question_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateQuestion(t *testing.T) {
	// Описываем тестовые кейсы
	tests := []struct {
		name                          string
		question                      string
		answer                        string
		countQuestionsByUserIdAtToday int
		hasAnswer                     bool
		isPremium                     bool
		expectedError                 bool
	}{
		{
			name:                          "Пользователь задал всего один вопрос и получает ответ",
			question:                      "Привет!",
			answer:                        "Добро пожаловать!",
			countQuestionsByUserIdAtToday: 1,
			hasAnswer:                     true,
			isPremium:                     false,
			expectedError:                 false,
		},
		{
			name:                          "Пользователь уже задал все возможные вопросы",
			question:                      "Привет!",
			countQuestionsByUserIdAtToday: 10,
			hasAnswer:                     false,
			isPremium:                     false,
			expectedError:                 true,
		},
		{
			name:                          "Премиум пользователь превысил лимит вопросов",
			question:                      "Привет!",
			countQuestionsByUserIdAtToday: 20,
			hasAnswer:                     false,
			isPremium:                     true,
			expectedError:                 true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Создаем моки
			repoMock := mocks.NewMockRepository(ctrl)
			premiumServiceMock := mocks.NewMockPremiumService(ctrl)
			answerClientMock := mocks.NewMockAnswerClient(ctrl)

			// Задаем поведение для checkLimit
			repoMock.EXPECT().
				CountQuestionsByUserIdAtToday(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(test.countQuestionsByUserIdAtToday, nil).
				Times(1)

			premiumServiceMock.EXPECT().
				CheckPremium(gomock.Any(), gomock.Any()).
				Return(test.isPremium, nil).
				Times(1)

			// Задаем поведение для Create и answerClient, если не ожидается ошибка
			if !test.expectedError {
				repoMock.
					EXPECT().
					Create(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entity.Question{}, nil).
					Times(1)

				if test.hasAnswer {
					answerClientMock.
						EXPECT().
						Request(test.question).
						Return([]string{test.answer}, nil).
						Times(1)
				}
			}

			// Создаем экземпляр сервиса с моками
			service := &Service{
				repo:                    repoMock,
				premiumService:          premiumServiceMock,
				answerClient:            answerClientMock,
				maxQuestionCount:        10,
				maxQuestionPremiumCount: 20,
			}

			// Вызываем метод сервиса
			answer, err := service.CreateQuestion(ctx, "00000000-0000-0000-0000-000000000001", test.question)

			if test.expectedError {
				assert.NoError(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, answer)
				if test.hasAnswer {
					assert.Equal(t, []string{test.answer}, answer.Texts)
				} else {
					assert.Contains(t, answer.Texts, "У вас превышен порог запросов за последние 24ч")
					assert.Contains(t, answer.Texts, "Обратитесь к @alexan_25")
				}
			}
		})
	}
}
