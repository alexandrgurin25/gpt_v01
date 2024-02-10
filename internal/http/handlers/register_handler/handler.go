package register_handler

import (
	"encoding/json"
	"net/http"

	"app/internal/service/register_service"
)

type handler struct {
	service *register_service.Service
}

func New(service *register_service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	var in RegisterDtoIn
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.Register(register_service.RegisterDto(in))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(struct{}{})
}
