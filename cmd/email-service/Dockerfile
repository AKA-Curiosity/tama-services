# Используем официальный образ Go с нужной версией
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum, чтобы установить зависимости
COPY go.mod go.sum ./

# Выполняем go mod tidy
RUN go mod tidy

# Копируем все исходные файлы из обеих директорий
COPY cmd/email-service ./cmd/email-service
COPY internal/email ./internal/email

# Копируем .env файл
COPY .env .env

# Переходим в директорию с main.go
WORKDIR /app/cmd/email-service

# Собираем приложение
RUN go build -o email-service .

# Финальный образ для продакшн
FROM debian:bullseye-slim

# Копируем бинарник из предыдущего шага
COPY --from=builder /app/cmd/email-service/email-service /usr/local/bin/email-service

# Копируем .env файл, если нужно, в финальный образ
COPY .env .env

# Устанавливаем рабочую директорию для контейнера
WORKDIR /app

# Запуск сервиса
CMD ["email-service"]

