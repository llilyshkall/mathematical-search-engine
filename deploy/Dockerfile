# Используем официальный образ Go
FROM golang:1.23.5-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/app

# Используем минимальный образ для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранное приложение из builder
COPY --from=builder /app/myapp .
COPY --from=builder /app/internal/templates /root/internal/templates

# Экспортируем порт, если это необходимо
EXPOSE 8080

# Запускаем приложение
CMD ["./myapp"]