package get_available_questions_count_handler

import (
	"encoding/json"
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
	userId, ok := r.Context().Value("userId").(string) // получение userId из контекста запроса

	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	count, err := h.service.AvailableCount(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out := &GetAvailableQuestionsCountOut{
		Count: count,
	}

	json.NewEncoder(w).Encode(out)
}
