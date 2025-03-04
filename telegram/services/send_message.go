package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// SendMessage sends a repository notification to a Telegram channel with an attached image
// It formats the message with HTML tags and includes the repository URL
func SendMessage(ctx context.Context, b *bot.Bot, url string, text string, fileData []byte) {
	username_repo := strings.TrimPrefix(url, "https://github.com/")
	telegramMessage := fmt.Sprintf("<a href=\"%s\">🔗 %s</a> %s\n\n<b><a href=\"https://t.me/github_ukraine\">🤖 GitHub Repositories</a></b>", url, username_repo, text)

	params := &bot.SendPhotoParams{
		ChatID:    os.Getenv("CHANNEL_ID"),
		Photo:     &models.InputFileUpload{Filename: "github.png", Data: bytes.NewReader(fileData)},
		Caption:   telegramMessage,
		ParseMode: models.ParseModeHTML,
	}

	if _, err := b.SendPhoto(ctx, params); err != nil {
		log.Printf("Error sending message: %v", err)
		return
	}
}
