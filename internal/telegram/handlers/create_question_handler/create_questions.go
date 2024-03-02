package create_question_telegram_handler

import (
	"app/internal/service/question_service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler struct {
	service *question_service.Service
}

func New(service *question_service.Service) *handler {
	return &handler{service: service}
}

func (h *handler) Handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	if update.Message == nil {
		return
	}

	answer, err := h.service.Create("5d0b8e8a-f38d-4767-ac03-ac171429da4e", update.Message.Text)

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
