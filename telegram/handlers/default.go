package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	commands := []models.BotCommand{
		{
			Command:     "start",
			Description: "Як finish тільки навпаки 🤷‍♀️",
		},
		{
			Command:     "help",
			Description: "Потрібна допомога? 🤔",
		},
	}

	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands:     commands,
		Scope:        nil,
		LanguageCode: "",
	})
}
