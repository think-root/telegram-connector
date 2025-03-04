package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// StartHandler responds to the /start command by sending a link to the GitHub Ukraine Telegram channel
func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "https://t.me/github_ukraine",
	})
}
