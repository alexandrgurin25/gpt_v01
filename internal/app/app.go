package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"app/internal/clients/gigachat"
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
	create_question_telegram_handler "app/internal/telegram/handlers/create_question_handler"
)

type app struct {
}

func New() *app {
	return &app{}
}

func (a *app) Start() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Создаем новое подключение к базе данных
	dataBase, err := database.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	userRepository := user_repository.New(dataBase)
	questionRepository := question_repository.New(dataBase)

	gigachat := gigachat.New()

	/*
		Пока не решена дальнейшая реализация 2 одновременно работающих
		нейронных сетей, поэтому пока используем наиболее простую (GigaChat)
	*/
	//gptchat := openai.New()
	
	registerService := register_service.New(userRepository)
	loginService := login_service.New(userRepository)
	questionService := question_service.New(questionRepository, gigachat)

	telegramHandler := create_question_telegram_handler.New(questionService)

	registerHandler := register_handler.New(registerService)
	loginHandler := login_handler.New(loginService)
	createQuestionHandler := create_question_handler.New(questionService)
	getAvailableQuestionsCountHandler := get_available_questions_count_handler.New(questionService)

	go func() {
		telegramChatApi, exists := os.LookupEnv("TELEGRAM_BOT_API")

		if !exists {
			log.Println("TelegramChatApi NOT FOUNT IN .env %w", err)
		}

		bot, err := tgbotapi.NewBotAPI(telegramChatApi)
		if err != nil {
			log.Println("Failed to create Telegram bot:", err)
			return
		}

		updateConfig := tgbotapi.NewUpdate(0)

		updateConfig.Timeout = 10

		updates := bot.GetUpdatesChan(updateConfig)

		for update := range updates {
			telegramHandler.Handle(update, bot)
		}
	}()

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
	log.Fatal(http.ListenAndServe(":8080", router))
}
