package handlers

import (
	"context"
	"log"
	"strings"

	"chappie_bot/config"
	"chappie_bot/helpers"
	"chappie_bot/repository"
	"chappie_bot/x"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendPostToXHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		repo, err := repository.GetRepository(1, true)
		if err != nil {
			log.Printf("Error getting repository: %v", err)
			return
		}

		if len(repo.Data.Items) == 0 {
			return
		}

		item := repo.Data.Items[0]
		username_repo := strings.TrimPrefix(item.URL, "https://github.com/")
		image_name := "./tmp/gh_project_img/image.png"

		err = repository.Socialify(username_repo)
		if err != nil {
			log.Println(err)
			err := helpers.CopyFile("./assets/banner.jpg", "./tmp/gh_project_img/image.png")
			if err != nil {
				log.Fatalf("Failed to copy file: %v", err)
			}
		}

		x_posted := x.CreateXPost(item.Text, item.URL, image_name)
		if x_posted {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text:   "✅ Message successfully sent to X",
			})
			log.Println("Message successfully sent to X")
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text:   "❌ Message not sent to X",
			})
			log.Println("Message not sent to X")
		}

		err = helpers.RemoveAllFilesInFolder("./tmp/gh_project_img")
		if err != nil {
			log.Println(err)
		}
	}
}
