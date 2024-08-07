package create_question_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/internal/service/question_service"
)

type handler struct {
	service *question_service.Service
}

func New(service *question_service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value("userId").(string)

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	var in CreateQuestionDtoIn
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateQuestion(ctx, userId, in.Text)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CreateQuestionDtoOut(*result))
}
