package middlewares

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// LoggingMiddleware logs incoming Telegram updates that contain messages
// It follows the middleware pattern by wrapping the next handler in the chain
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
		// Continue to the next handler in the chain
		next(ctx, b, update)
	}
}
