package common

import (
	"errors"
	"log"
	"net/http"
)

// HandleHttpError обрабатывает ошибки HTTP и отправляет соответствующий ответ клиенту.
func HandleHttpError(w http.ResponseWriter, err error) {
	log.Printf("%v", err)

	if errors.Is(err, NotFoundError) {
		// Если ошибка является NotFoundError, отправляем статус "404 Not Found" клиенту.
		http.Error(w, http.StatusText(404), 404)
		return
	}

	if errors.Is(err, LogicError) {
		// Если ошибка является LogicError, отправляем статус "400 Bad Request" клиенту.
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if errors.Is(err, ForbiddenError) {
		// Если ошибка является ForbiddenError, отправляем статус "401 Unauthorized" клиенту.
		http.Error(w, http.StatusText(401), 401)
		return
	}

	// Если ошибка не соответствует ни одному из вышеперечисленных типов, отправляем статус "500 Internal Server Error" клиенту.
	http.Error(w, http.StatusText(500), 500)
}
