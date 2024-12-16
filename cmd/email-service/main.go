package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"tama-services/internal/email"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Читаем настройки SMTP
	config := email.SMTPConfig{
		Server:   os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}

	// Создаём сервис для отправки email
	emailService := email.NewService(config)

	// HTTP-обработчик для отправки писем
	http.HandleFunc("/send-email", emailService.SendEmail)

	log.Println("Сервис запущен на порту 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
