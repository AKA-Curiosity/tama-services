package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
)

// SMTP настройки
const (
	SMTPServer = "mail.ta-ma.ru"
	SMTPPort   = 465
	Email      = "no-reply@ta-ma.ru"
	Password   = "vH7xU9jM1qxH4gX4"
)

// sendEmail отправляет письмо через SMTP
func sendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", Email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(SMTPServer, SMTPPort, Email, Password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // Отключаем проверку сертификата
	}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("ошибка отправки письма: %w", err)
	}
	return nil
}

// EmailHandler обрабатывает POST-запросы для отправки писем
func EmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST-запросы", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем параметры из тела запроса
	to := r.URL.Query().Get("to")
	if to == "" {
		http.Error(w, "Отсутствует email-получатель", http.StatusBadRequest)
		return
	}

	subject := "Код подтверждения"
	body := "Ваш код подтверждения: 123456"

	// Отправляем письмо
	err := sendEmail(to, subject, body)
	if err != nil {
		http.Error(w, "Ошибка при отправке письма: "+err.Error(), http.StatusInternalServerError)
		log.Println("Ошибка отправки:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"Письмо отправлено на %s"}`, to)
}
