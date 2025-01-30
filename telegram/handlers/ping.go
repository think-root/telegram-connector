package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func PingHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   "pong 🏓",
	})
}
