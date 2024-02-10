package create_question_handler

type CreateQuestionDtoIn struct {
	Text string `json:"text"`
}

type CreateQuestionDtoOut struct {
	Text string `json:"text"`
}
