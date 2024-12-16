# Используем официальный образ Go с нужной версией
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию для сборки
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Выполняем go mod tidy для установки зависимостей
RUN go mod tidy

# Создаём каталог vendor и сохраняем все зависимости
RUN go mod vendor

# Копируем исходники проекта
COPY cmd ./cmd
COPY internal ./internal

# Переходим в директорию с main.go для сборки
WORKDIR /app/cmd/email-service

# Собираем приложение
RUN go build -o email-service .

# Финальный образ для продакшн
FROM debian:bullseye-slim

# Копируем бинарник из предыдущего шага
COPY --from=builder /app/cmd/email-service/email-service /usr/local/bin/email-service

# Копируем .env файл, если нужно, в финальный образ
COPY cmd/email-service/.env .env

# Устанавливаем рабочую директорию для контейнера
WORKDIR /app

# Если нужно использовать переменные окружения из .env, можно сделать это через запуск скрипта
# Например, используем envsubst для замены значений переменных
RUN apt-get update && apt-get install -y gettext-base
RUN envsubst < .env > /etc/environment

# Запуск сервиса
CMD ["email-service"]
