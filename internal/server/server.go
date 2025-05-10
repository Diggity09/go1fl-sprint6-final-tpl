package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// Функция для создания нового сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем новый роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем HTTP сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Создаем и возвращаем наш сервер
	return &Server{
		Logger: logger,
		Server: httpServer,
	}
}

// Метод для запуска сервера
func (s *Server) Start() error {
	s.Logger.Println("Запуск сервера на порту :8080")
	return s.Server.ListenAndServe()
}
