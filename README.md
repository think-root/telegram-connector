# telegram-bridge

[![Go Version](https://img.shields.io/github/go-mod/go-version/think-root/telegram-bridge)](https://github.com/think-root/telegram-bridge)
[![License](https://img.shields.io/github/license/think-root/telegram-bridge)](LICENSE)
[![Version](https://img.shields.io/github/v/release/think-root/telegram-bridge)](https://github.com/think-root/telegram-bridge/releases)
[![Changelog](https://img.shields.io/badge/changelog-view-blue)](CHANGELOG.md)
[![Deploy Status](https://github.com/think-root/telegram-bridge/workflows/Deploy%20telegram-bridge/badge.svg)](https://github.com/think-root/telegram-bridge/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/think-root/telegram-bridge)](https://goreportcard.com/report/github.com/think-root/telegram-bridge)

An AI-powered social media management bot that automatically publishes curated content across multiple platforms including Telegram, Twitter, and WhatsApp. Built to work with [content-alchemist](https://github.com/think-root/content-alchemist).

[Read the origin story (in Ukrainian): [Part 1](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-1-VywRW) | [Part 2](https://drukarnia.com.ua/articles/yak-chatgpt-vede-za-mene-kanal-v-telegram-i-u-nogo-ce-maizhe-vikhodit-chastina-2-X9Yjz)]

## Features

- 🔄 Multi-platform posting (Telegram, Twitter, WhatsApp)
- 🤖 AI-powered content curation
- ⚡ Automated GitHub repository discovery
- 📊 Post scheduling and management
- 🎯 Custom bot commands

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/think-root/telegram-bridge.git
cd telegram-bridge
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

| Variable | Description |
|----------|-------------|
| BOT_TOKEN | Telegram bot token from [@BotFather](https://core.telegram.org/bots) |
| ADMIN_ID | Your Telegram user ID |
| CHANNEL_ID | Target Telegram channel ID |
| CHAPPIE_SERVER_URL | URL of your content-alchemist instance |
| CHAPPIE_SERVER_BEARER | Authentication token for content-alchemist |
| X_API_KEY | API key for X (Twitter) integration |
| X_URL | URL for self-hosted X API server |

### Optional Environment Variables

⚠️ **Warning**: WhatsApp integration is unofficial and may risk account suspension

| Variable | Description | Default |
|----------|-------------|---------|
| ENABLE_CRON | Enable post collection cron job | false |
| WAPP_ENABLE | Enable WhatsApp integration | false |
| WAPP_JID | WhatsApp chat JID | - |
| WAPP_TOKEN | WhatsApp authentication token | - |
| WAPP_SERVER_URL | WhatsApp server URL | - |

## Bot Commands

| Command | Description |
|---------|-------------|
| /add [urls] | Add GitHub repositories to content queue |
| /next [count] | Preview upcoming posts |
| /gen [count] | Generate posts from GitHub trends |
| /info | View database statistics |
| /xsend | Send post to X (Twitter) |

## Dependencies

- [Docker](https://docs.docker.com/engine/install/)
- [content-alchemist](https://github.com/think-root/content-alchemist)
- [Docker Compose](https://docs.docker.com/compose/install/) (optional)
- [X API Server](https://github.com/think-root/x) (optional, for Twitter integration)
- [WhatsApp Server](https://github.com/think-root/wapp) (optional, unofficial and may risk account suspension)

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
go run ./cmd/bot/main.go

# Build binary
go build -o telegram-bridge ./cmd/bot/main.go
```

## License

This project is licensed under the BSD 2-Clause License - see the [LICENSE](LICENSE) file for details.