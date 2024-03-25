package add_user_handler

import (
	"app/internal/service/telegram_user_service"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler struct {
	service *telegram_user_service.Service
}

func New(service *telegram_user_service.Service) *handler {
	return &handler{service: service}
}

func (h *handler) Handle(update tgbotapi.Update, bot *tgbotapi.BotAPI) {

	_, err := h.service.CreateUserIdByChatId(update.Message.Chat.ID)

	if err != nil {
		log.Println("add_user_handler.Handle couldn't CreateUserIdByChatId->", err)
	}

	messageCommandStart := fmt.Sprintf("Добро пожаловать, %s! Я - нейросетевая языковая модель, созданная командой @alexan_25."+
		"Я могу помочь вам с различными задачами, такими как написание текстов, ответы на вопросы, перевод с одного языка на"+
		" другой и многое другое.  Если у вас есть какие-либо вопросы или запросы, не стесняйтесь обращаться ко мне.",
		update.Message.Chat.FirstName)

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, messageCommandStart))

	if err != nil {
		log.Println("add_user_handler.Handle bot couldn't send messageCommandStart", err)
	}
}
