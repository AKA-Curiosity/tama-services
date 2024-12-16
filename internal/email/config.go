package email

// SMTPConfig хранит данные для подключения к SMTP-серверу
type SMTPConfig struct {
	Server   string
	Port     string
	Email    string
	Password string
}
