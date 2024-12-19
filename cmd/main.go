package main

import (
	"log"
	"net/http"
	"tama-services/internal/auth"
	"tama-services/internal/db"
	"tama-services/internal/email"
)

func main() {
	// Инициализация MongoDB
	db.InitMongoDB("mongodb://admin:Bo5aK5!t@212.67.11.16:27017/admin")

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/send-email", email.EmailHandler)

	log.Println("Starting server on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
