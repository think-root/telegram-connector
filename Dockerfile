# Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG APP_VERSION=dev
RUN go build -ldflags="-X 'telegram-bridge/config.APP_VERSION=${APP_VERSION}'" -o telegram-bridge ./cmd/bot/main.go

# Runtime
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/telegram-bridge .
COPY .env /app/.env
COPY assets/ /app/assets/
CMD ["./telegram-bridge"]