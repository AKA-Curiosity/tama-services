package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"tama-services/internal/db"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
	Error  string `json:"error,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Проверка пользователя через MongoDB
	user, err := db.FindUserByEmail(req.Email)
	if err != nil {
		http.Error(w, `{"error":"Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	log.Printf("Found user: %+v", user) // Логирование пользователя

	// Сравнение пароля с хешом в базе
	if user.Password != req.Password {
		http.Error(w, `{"error":"Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	// Генерация токена (например, JWT)
	token := "your_generated_jwt_token" // Замени на реальную генерацию токена

	resp := LoginResponse{
		Status: "success",
		Token:  token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
