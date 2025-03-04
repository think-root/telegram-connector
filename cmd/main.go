package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"telegram-connector/server"
	"telegram-connector/telegram/handlers"
	"telegram-connector/telegram/middlewares"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

// init loads environment variables from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		os.Exit(1)
	}
}

func main() {
	// Set application version from env or use default
	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "development"
	}

	// Setup context with cancellation on interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Configure bot options with middleware
	opts := []bot.Option{
		bot.WithMiddlewares(middlewares.LoggingMiddleware),
	}

	// Initialize Telegram bot
	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		log.Println(err.Error())
	}

	// Configure and start HTTP server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	go server.Start(serverPort, ctx, b)

	// Register bot command handlers
	handlers.RegisterCommands(b)

	log.Printf("bot: running\n")
	log.Printf("app version: %s\n", appVersion)

	// Start bot and wait for context cancellation
	b.Start(ctx)
}
