package handlers

import (
	"bytes"
	"chappie_bot/config"
	"chappie_bot/helpers"
	"chappie_bot/repository"
	"chappie_bot/whatsapp"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendPostHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		repo, err := repository.GetRepository(1, false)
		if err != nil {
			log.Printf("Error getting repository: %v", err)
			return
		}

		if len(repo.Data.Items) == 0 {
			return
		}

		item := repo.Data.Items[0]
		username_repo := strings.TrimPrefix(item.URL, "https://github.com/")
		message := fmt.Sprintf("<a href=\"%s\">🔗 %s</a> %s", item.URL, username_repo, item.Text)

		err = repository.Socialify(username_repo)
		if err != nil {
			log.Println(err)
			err := helpers.CopyFile("./assets/github_octopus_logo.png", "./tmp/gh_project_img/image.png")
			if err != nil {
				log.Fatalf("Failed to copy file: %v", err)
			}
		}

		fileData, errReadFile := os.ReadFile("./tmp/gh_project_img/image.png")
		if errReadFile != nil {
			log.Printf("error reading file: %v\n", errReadFile)
			return
		}

		params := &bot.SendPhotoParams{
			ChatID: config.CHANNEL_ID,
			Photo:  &models.InputFileUpload{Filename: "github.png", Data: bytes.NewReader(fileData)},
			Caption: message +
				"\n\n<b><a href=\"https://t.me/github_ukraine\">🤖 GitHub Repositories</a></b>",
			ParseMode: models.ParseModeHTML,
		}

		if _, err := b.SendPhoto(ctx, params); err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}

		err = helpers.RemoveAllFilesInFolder("./tmp/gh_project_img")
		if err != nil {
			log.Println(err)
		}

		wapp := whatsapp.SendMessageToWhatsApp(message, config.WAPP_JID)
		if wapp {
			log.Println("Message successfully sent to whatsapp")
		} else {
			log.Println("Message not sent to whatsapp")
		}

		if result, err := repository.UpdateRepositoryPosted(item.URL, true); err != nil {
			log.Printf("Error updating repository posted status: %v", err)
		} else if result {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: config.ADMIN_ID,
				Text:   "Повідомлення успішно відправлене",
			})
		}
	}
}
