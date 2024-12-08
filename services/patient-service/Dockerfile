# Используем официальный образ Golang как базовый
FROM golang:1.23.4-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Скопируем весь монорепозиторий (для go.mod и кода)
COPY . .

# Переходим в директорию конкретного микросервиса
WORKDIR /app/services/service1 ./services/patient-service

# Загрузка модулей
RUN go mod tidy

# Сборка приложения
RUN go build -o app ./cmd/service-name # Замените на точку входа микросервиса

# Минимальный образ для запуска
FROM debian:bullseye-slim

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/services/service1/app .

# Указываем порт
EXPOSE 8080 # Замените на нужный порт

# Запуск приложения
CMD ["./app"]