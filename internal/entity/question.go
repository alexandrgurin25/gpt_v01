package entity

type Question struct {
	ID        int64  `json:"id"`
	UserId    string `json:"user_id"` // идентификатор пользователя задавшего вопрос
	Text      string `json:"text"`    // текст запроса
	CreatedAt string `json:"created_at"`
}
