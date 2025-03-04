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
ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}
ENV PORT=9111
COPY --from=builder /app/telegram-connector .
COPY .env /app/.env
EXPOSE ${PORT}
CMD ["./telegram-connector"]