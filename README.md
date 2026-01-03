# telegram-connector

This project is part of the [content-maestro](https://github.com/think-root/content-maestro) repository. If you want Telegram integration and automatic publishing of posts there as well, you need to deploy this app.

## Description

A Go-based HTTP service that receives publishing requests and relays them to Telegram via the Bot API. It protects every request with API key middleware, accepts multipart payloads with optional media, and focuses on a lightweight local workflow.

## Prerequisites

- Go 1.21+ (for local development)
- A Telegram bot token from [@BotFather](https://core.telegram.org/bots)

## Setup

1. **Clone the repository:**

2. **Create and configure the environment file:**

   ```bash
   cp .env.example .env
   ```

   Fill in the required variables:

   ```
   BOT_TOKEN=your_telegram_bot_token
   CHANNEL_ID=your_target_channel_id
   X_API_KEY=your_private_api_key
   SERVER_PORT=8080
   ```

3. **Install Go dependencies (optional for local runs):**

   ```bash
   go mod download
   ```

4. **Run the service:**

   ```bash
   go run ./cmd/main.go
   ```

   To build a binary for distribution:

   ```bash
   go build -o telegram-connector ./cmd/main.go
   ```

## API

All endpoints require the `X-API-Key` header for authentication.

### Authentication

| Header      | Type   | Required | Description                          |
|-------------|--------|----------|--------------------------------------|
| `X-API-Key` | string | Yes      | API key defined in your `.env` file. |

**Error Response (401 Unauthorized):**

```json
{
  "detail": "Invalid or missing API key"
}
```

---

### POST `/telegram/send-message`

Publishes a post to the configured Telegram channel. Supports text, optional image attachment, and optional URL.

#### Request

**Content-Type:** `multipart/form-data`

| Parameter | Type   | Required | Description                                                        |
|-----------|--------|----------|--------------------------------------------------------------------|
| `text`    | string | Yes      | Message body delivered to the channel.                             |
| `url`     | string | No       | URL appended to the post.                                          |
| `image`   | file   | No       | Image file uploaded and attached to the Telegram message.          |

#### Examples

**Simple post:**

```bash
curl -X POST "http://localhost:8080/telegram/send-message" \
  -H "X-API-Key: your_api_key" \
  -F "text=Hello, Telegram!"
```

**Post with image and URL:**

```bash
curl -X POST "http://localhost:8080/telegram/send-message" \
  -H "X-API-Key: your_api_key" \
  -F "text=Fresh article drop" \
  -F "url=https://example.com/post" \
  -F "image=@/path/to/image.jpg"
```

#### Response

**Success (200 OK):**

```json
{
  "status": "OK",
  "message": "Message sent successfully"
}
```

**Error:**

```json
{
  "error": "Error message describing what went wrong"
}
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
