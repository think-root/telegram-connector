# telegram-connector

[![Go Version](https://img.shields.io/github/go-mod/go-version/think-root/telegram-connector)](https://github.com/think-root/telegram-connector)
[![License](https://img.shields.io/github/license/think-root/telegram-connector)](LICENSE)
[![Version](https://img.shields.io/github/v/release/think-root/telegram-connector)](https://github.com/think-root/telegram-connector/releases)
[![Changelog](https://img.shields.io/badge/changelog-view-blue)](CHANGELOG.md)
[![Deploy Status](https://github.com/think-root/telegram-connector/workflows/Deploy%20telegram-connector/badge.svg)](https://github.com/think-root/telegram-connector/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/think-root/telegram-connector)](https://goreportcard.com/report/github.com/think-root/telegram-connector)

This app is part of [content-alchemist](https://github.com/think-root/content-alchemist). It is a Telegram integration that performs a simple function — publishing content to a Telegram [channel](https://t.me/github_ukraine).

## Quick Start

1. Clone the repository:

```bash
git clone https://github.com/think-root/telegram-connector.git
cd telegram-connector
```

2. Configure environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Deploy with Docker:

```bash
docker compose up -d
```

## Detailed Configuration

### Required Environment Variables

| Variable    | Description                                                          |
| ----------- | -------------------------------------------------------------------- |
| BOT_TOKEN   | Telegram bot token from [@BotFather](https://core.telegram.org/bots) |
| CHANNEL_ID  | Target Telegram channel ID                                           |
| X_API_KEY   | Your key for API protection                                          |
| SERVER_PORT | Port for the HTTP server                                             |

## Dependencies

- [content-alchemist](https://github.com/think-root/content-alchemist)
- [content-maestro](https://github.com/think-root/content-maestro)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/) (optional)

## API Endpoints

### Send Message to Telegram Channel

```
POST /telegram/send-message
```

Headers:

- `X-API-Key`: Your API key (from X_API_KEY env var)
- `Content-Type`: multipart/form-data

Form Data:

- `url`: URL to be included in the message
- `text`: Text message to send to the channel
- `image`: Image file to attach to the message

Example Response:

```json
{
  "status": "OK",
  "message": "Message sent successfully"
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup

```bash
# Install Go
go mod download

# Run locally
go run ./cmd/main.go

# Build binary
go build -o telegram-connector ./cmd/main.go
```
