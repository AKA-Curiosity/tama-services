package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tama-services/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Обработчик логина
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	// Проверяем пользователя в базе
	if !CheckUserCredentials(creds.Username, creds.Password) {
		http.Error(w, "Неверные имя пользователя или пароль", http.StatusUnauthorized)
		return
	}

	// Если проверка прошла, возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Вход успешен! Добро пожаловать, %s", creds.Username)
}

// Функция для проверки логина и пароля
func CheckUserCredentials(username, password string) bool {
	collection := db.GetUserCollection()
	filter := bson.M{"username": username, "password": password}

	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false
	}

	return true
}
