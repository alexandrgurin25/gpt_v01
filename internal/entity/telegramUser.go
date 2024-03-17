package entity

type TelegramUser struct {
	UserId string `json:"user_id"`
	ChatId int64    `json:"chat_id"`
}
