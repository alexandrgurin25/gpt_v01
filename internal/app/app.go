package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"app/internal/database"
	"app/internal/http/handlers/create_question_handler"
	"app/internal/http/handlers/get_available_questions_count_handler"
	"app/internal/http/handlers/login_handler"
	"app/internal/http/handlers/register_handler"
	"app/internal/http/middlewares"
	"app/internal/repository/question_repository"
	"app/internal/repository/user_repository"
	"app/internal/service/login_service"
	"app/internal/service/question_service"
	"app/internal/service/register_service"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	// Создаем новое подключение к базе данных
	dataBase, err := database.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	userRepository := user_repository.New(dataBase)
	questionRepository := question_repository.New(dataBase)

	loginService := login_service.New(userRepository)
	registerService := register_service.New(userRepository)
	questionService := question_service.New(questionRepository)

	loginHandler := login_handler.New(loginService)
	registerHandler := register_handler.New(registerService)
	createQuestionHandler := create_question_handler.New(questionService)
	getAvailableQuestionsCountHandler := get_available_questions_count_handler.New(questionService)

	// Создаем новый роутер Chi
	router := chi.NewRouter()

	// Используем промежуточное ПО для обработки запросов
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	router.Route("/auth", func(router chi.Router) {
		router.Post("/login", loginHandler.Handle)
		router.Post("/register", registerHandler.Handle)
	})

	router.Route("/questions", func(router chi.Router) {
		router.
			With(middlewares.AuthMiddleware).
			Post("/", createQuestionHandler.Handle)

		router.
			With(middlewares.AuthMiddleware).
			Get("/available-count", getAvailableQuestionsCountHandler.Handle)
	})

	// Запуск HTTP-сервера и обработка запросов с помощью роутера
	log.Fatal(http.ListenAndServe(":80", router))
}
