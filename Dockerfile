# Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG APP_VERSION=dev
RUN go build -ldflags="-X 'telegram-connector/config.APP_VERSION=${APP_VERSION}'" -o telegram-connector ./cmd/main.go

# Runtime
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/telegram-connector .
COPY .env /app/.env
CMD ["./telegram-connector"]