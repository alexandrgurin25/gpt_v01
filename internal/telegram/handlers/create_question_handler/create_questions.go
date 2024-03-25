package create_question_telegram_handler

import (
	"app/internal/service/question_service"
	"app/internal/service/telegram_user_service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler struct {
	questionService *question_service.Service
	telegramService *telegram_user_service.Service
}

func New(qService *question_service.Service, tService *telegram_user_service.Service) *handler {
	return &handler{questionService: qService, telegramService: tService}
}

func (h *handler) Handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	if update.Message == nil {
		return
	}

	user, err := h.telegramService.CreateUserIdByChatId(update.Message.Chat.ID)
	if err != nil {
		log.Println("", err)
		return
	}

	answer, err := h.questionService.Create(user.UserId, update.Message.Text)

	if err != nil {
		log.Println("", err)
		return
	}

	for _, response := range answer.Texts {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			if _, err := bot.Send(msg); err != nil {
				log.Println("Failed to send message via Telegram bot:", err)

			}
		}
	}

}
	