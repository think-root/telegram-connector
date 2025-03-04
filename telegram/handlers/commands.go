package handlers

import (
	"github.com/go-telegram/bot"
)

// RegisterCommands registers all available bot commands with their respective handlers
func RegisterCommands(b *bot.Bot) {
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/start",
		bot.MatchTypeExact,
		StartHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/ping",
		bot.MatchTypePrefix,
		PingHandler,
	)
}
