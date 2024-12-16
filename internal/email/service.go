package email

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"net/http"
	"net/mail"
	"net/smtp"
)

// Service отвечает за отправку email
type Service struct {
	config SMTPConfig
}

// NewService создаёт новый email-сервис
func NewService(config SMTPConfig) *Service {
	return &Service{config: config}
}

// EmailRequest структура для парсинга данных из запроса
type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// isValidEmail проверяет корректность email-адреса
func isValidEmail(emailAddr string) bool {
	_, err := mail.ParseAddress(emailAddr)
	return err == nil
}

// SendEmail отправляет email на основе данных из HTTP-запроса
func (s *Service) SendEmail(w http.ResponseWriter, r *http.Request) {
	// Проверка метода
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим JSON-данные из тела запроса
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка парсинга запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Проверка обязательных полей
	if req.To == "" || req.Subject == "" || req.Body == "" {
		http.Error(w, "Поле 'to', 'subject' и 'body' обязательны", http.StatusBadRequest)
		return
	}

	// Валидация email-адреса
	if !isValidEmail(req.To) {
		http.Error(w, "Некорректный email адрес", http.StatusBadRequest)
		return
	}

	// Создание письма
	e := email.NewEmail()
	e.From = s.config.Email
	e.To = []string{req.To}
	e.Subject = req.Subject
	e.Text = []byte(req.Body)

	// Настройка SMTP-клиента
	auth := smtp.PlainAuth("", s.config.Email, s.config.Password, s.config.Server)
	err := e.SendWithTLS(fmt.Sprintf("%s:%s", s.config.Server, s.config.Port), auth, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.config.Server,
	})
	if err != nil {
		http.Error(w, "Ошибка отправки письма: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Письмо успешно отправлено!"))
}
