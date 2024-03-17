package create_question_telegram_handler

import (
	"app/internal/service/question_service"
	"app/internal/service/telegram_user_service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler struct {
	question_service *question_service.Service
	telegram_service *telegram_user_service.Service
}

func New(qService *question_service.Service, tService *telegram_user_service.Service) *handler {
	return &handler{question_service: qService, telegram_service: tService}
}

func (h *handler) Handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	if update.Message == nil {
		return
	}

	user, err := h.telegram_service.CreateFieldUserId(update.Message.Chat.ID)
	if err != nil {
		log.Println("", err)
		return
	}

	answer, err := h.question_service.Create(user.UserId, update.Message.Text)

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
