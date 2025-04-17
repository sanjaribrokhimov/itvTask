# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Установка зависимостей для сборки
RUN apk add --no-cache git

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Установка необходимых пакетов
RUN apk --no-cache add ca-certificates tzdata

# Копируем бинарный файл из builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Создаем непривилегированного пользователя
RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8080

CMD ["./main"] 