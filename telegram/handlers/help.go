package handlers

import (
	"chappie_bot/config"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    userID,
			ParseMode: models.ParseModeHTML,
			Text:      "<b>/add</b> - можна написати перелік (або один) GitHub репо через пробіл, для них буде згенеровано короткий опис, і вони будуть додані в БД\n\n<b>/next</b> - показує наступний пост який буде опубліковано в каналі, а якщо написати цифру через пробіл то покаже відповідну кількість постів\n\n<b>/gen</b> - та шо, генерує пост для каналу з трендів GitHub, якщо поставити цифру через пробіл можна вказати кількість\n\n<b>/info</b> - показати інфу по постам з БД\n\n<b>app version:</b> <i>" + config.APP_VERSION + "</i>",
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:             userID,
			ParseMode:          models.ParseModeHTML,
			LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: bot.True()},
			Text:               "Не знаю чим я можу вам допомогти 💁‍♀️\n\nAле от ✨<b>ВИ</b>✨ так, саме ✨<b>ВИ</b>✨ можете <b><a href=\"https://send.monobank.ua/jar/dzBdJ3737\">ДОПОМОГТИ</a></b> зменшити кількість загарбників <b>України</b> 🇺🇦",
		})
	}
}
