# Используем официальный образ Go с нужной версией
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Выполняем go mod tidy
RUN go mod tidy

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN go build -o email-service .

# Финальный образ для продакшн
FROM debian:bullseye-slim

# Копируем бинарник из предыдущего шага
COPY --from=builder /app/email-service /usr/local/bin/email-service

# Устанавливаем рабочую директорию для контейнера
WORKDIR /app

# Запуск сервиса
CMD ["email-service"]
