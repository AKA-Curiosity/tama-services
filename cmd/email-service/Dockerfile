# Используем официальный Go-образ для сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем все исходники проекта
COPY . .

# Сборка бинарного файла
RUN go build -o email-service ./cmd/email-service

# Финальный образ с минимальной загрузкой
FROM ubuntu:22.04

WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Копируем собранный бинарник
COPY --from=builder /app/email-service /usr/local/bin/email-service

# Открываем порт для сервиса
EXPOSE 8081

CMD ["email-service"]
