# Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o telegram-connector ./cmd/main.go

# Runtime
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/telegram-connector .
COPY .env /app/.env
ENV APP_VERSION=${APP_VERSION}
CMD ["./telegram-connector"]