package middlewares

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LoggingMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			jsonBytes, err := json.Marshal(update)
			if err == nil {
				jsonString := string(jsonBytes)
				log.Printf("%s\n\n", jsonString)
			} else {
				log.Printf("Error while marshalling update: %v\n", err)
			}
		}
		next(ctx, b, update)
	}
}
