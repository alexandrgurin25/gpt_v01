package login_handler

import (
	"encoding/json"
	"net/http"

	"app/internal/service/login_service"
)

type handler struct {
	service *login_service.Service
}

func New(service *login_service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	var in LoginDtoIn
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authDto, err := h.service.Login(login_service.LoginDto(in))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(LoginDtoOut(*authDto))
}
