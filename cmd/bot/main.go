package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"telegram-bridge/config"
	"telegram-bridge/cron"
	"telegram-bridge/helpers"
	"telegram-bridge/telegram/handlers"
	"telegram-bridge/telegram/middlewares"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(middlewares.LoggingMiddleware),
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}

	helpers.CreateDirIfNotExist("tmp/gh_project_img")

	b, err := bot.New(config.BOT_TOKEN, opts...)
	if err != nil {
		log.Println(err.Error())
	}

	handlers.RegisterCommands(b)

	cron.SendMessageCron(ctx, b)

	if config.ENABLE_CRON {
		log.Println("Starting CollectPostsCron...")
		cron.CollectPostsCron(ctx, b)
	}

	log.Printf("Bot successfully started ^_^ (app version: %s)\n\n", config.APP_VERSION)

	b.Start(ctx)
}
