# Этап сборки
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и .env (если нужно) и другие исходники
COPY go.mod /app/go.mod
COPY cmd/email-service/.env /app/.env
COPY cmd/email-service /app/cmd/email-service
COPY internal /app/internal

# Загружаем зависимости (перед сборкой)
RUN go mod tidy
RUN go mod download

# Сборка проекта
WORKDIR /app/cmd/email-service
RUN go build -o email-service .

# Финальный образ
FROM gcr.io/distroless/base:latest

# Копируем скомпилированное приложение в финальный образ
COPY --from=builder /app/cmd/email-service/email-service /email-service

# Вставляем .env, если нужно
COPY cmd/email-service/.env .env

# Устанавливаем команду для запуска приложения
ENTRYPOINT ["/email-service"]
