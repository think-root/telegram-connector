package handlers

import (
	"github.com/go-telegram/bot"
)

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
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/gen",
		bot.MatchTypePrefix,
		GeneratePostsHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/info",
		bot.MatchTypePrefix,
		DBStatisticHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/help",
		bot.MatchTypePrefix,
		HelpHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/add",
		bot.MatchTypePrefix,
		ManualAddPostHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/next",
		bot.MatchTypePrefix,
		NextPostHandler,
	)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/send",
		bot.MatchTypePrefix,
		SendPostHandler,
	)
}
