package handlers

import (
	"context"

	"telegram-connector/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    userID,
			ParseMode: models.ParseModeHTML,
			Text:      "<b>/add</b> - you can write a list (or one) GitHub repo separated by space, a short description will be generated for them, and they will be added to the DB\n\n<b>/next</b> - shows the next post that will be published in the channel, and if you write a number after space it will show the corresponding number of posts\n\n<b>/gen</b> - well yeah, generates a post for the channel from GitHub trends, if you put a number after space you can specify the quantity\n\n<b>/info</b> - show info about posts from DB\n\n<b>app version:</b> <i>" + config.APP_VERSION + "</i>",
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:             userID,
			ParseMode:          models.ParseModeHTML,
			LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: bot.True()},
			Text:               "I don't know how I can help you 💁‍♀️\n\nBut ✨<b>YOU</b>✨ yes, exactly ✨<b>YOU</b>✨ can <b><a href=\"https://send.monobank.ua/jar/dzBdJ3737\">HELP</a></b> reduce the number of russian invaders in <b>Ukraine</b> 🇺🇦",
		})
	}
}
