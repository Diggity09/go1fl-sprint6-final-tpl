package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "[MORSE SERVER] ", log.LstdFlags)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	err := srv.Start()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
