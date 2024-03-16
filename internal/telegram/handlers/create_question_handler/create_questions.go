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
// Получить из базы данных получить userId по chatId
// Метод репозитория getUserByChatId

	answer, err := h.service.Create("adfc5c8f-c94e-4326-9836-4eea1431412c", update.Message.Text)

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
