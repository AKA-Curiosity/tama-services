package main

import (
	"fmt"
	"log"
	"net/http"
	"tama-services/internal/auth"
	"tama-services/internal/db"
)

func main() {
	// Подключаемся к базе данных
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Настроим HTTP маршруты
	http.HandleFunc("/login", auth.LoginHandler)

	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
