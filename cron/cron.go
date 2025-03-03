package cron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"telegram-connector/config"
	"telegram-connector/helpers"
	"telegram-connector/repository"
	"telegram-connector/whatsapp"
	"telegram-connector/x"

	"github.com/go-co-op/gocron"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type generateRequest struct {
	MaxRepos           int    `json:"max_repos"`
	Since              string `json:"since"`
	SpokenLanguageCode string `json:"spoken_language_code"`
}

type generateResponse struct {
	Status    string   `json:"status"`
	Added     []string `json:"added"`
	DontAdded []string `json:"dont_added"`
}

func SendMessageCron(ctx context.Context, b *bot.Bot) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("10:12:00").SingletonMode().Do(func() {
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
		telegramMessage := fmt.Sprintf("<a href=\"%s\">🔗 %s</a> %s\n\n<b><a href=\"https://t.me/github_ukraine\">🤖 GitHub Repositories</a></b>", item.URL, username_repo, item.Text)

		image_name := "./tmp/gh_project_img/image.png"

		err = repository.Socialify(username_repo)
		if err != nil {
			log.Println(err)
			err := helpers.CopyFile("./assets/banner.jpg", image_name)
			if err != nil {
				log.Printf("Failed to copy file: %v", err)
				return
			}
		}

		fileData, errReadFile := os.ReadFile(image_name)
		if errReadFile != nil {
			log.Printf("error reading file: %v\n", errReadFile)
			return
		}

		params := &bot.SendPhotoParams{
			ChatID:    config.CHANNEL_ID,
			Photo:     &models.InputFileUpload{Filename: "github.png", Data: bytes.NewReader(fileData)},
			Caption:   telegramMessage,
			ParseMode: models.ParseModeHTML,
		}

		if _, err := b.SendPhoto(ctx, params); err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}

		if config.WAPP_ENABLE {
			whatsappMessage := fmt.Sprintf("🔗 %s\n\n%s\n\n🤖 GitHub Repositories", item.URL, item.Text)
			wapp := whatsapp.SendMessageToWhatsApp(whatsappMessage, config.WAPP_JID)
			if wapp {
				log.Println("Message successfully sent to whatsapp")
			} else {
				log.Println("Message not sent to whatsapp")
			}
		}

		x_posted := x.CreateXPost(item.Text, item.URL, image_name)
		if x_posted {
			log.Println("Message successfully sent to X")
		} else {
			log.Println("Message not sent to X")
		}

		if result, err := repository.UpdateRepositoryPosted(item.URL, true); err != nil {
			log.Printf("Error updating repository posted status: %v", err)
		} else if result {
			log.Println("Message successfully sent to the telegram channel")
		}

		err = helpers.RemoveAllFilesInFolder("./tmp/gh_project_img")
		if err != nil {
			log.Println(err)
		}
	})
	s.StartAsync()
}

func CollectPostsCron(ctx context.Context, b *bot.Bot) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("13:00:00").SingletonMode().Do(func() {
		log.Println("Collecting posts...")

		payload := generateRequest{
			MaxRepos:           1,
			Since:              "daily",
			SpokenLanguageCode: "en",
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Error marshaling request: %v", err)
			return
		}

		req, err := http.NewRequest("POST", config.CONTENT_ALCHEMIST_URL+"think-root/api/auto-generate/", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.CONTENT_ALCHEMIST_BEARER)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response: %v", err)
			return
		}

		var response generateResponse
		if err := json.Unmarshal(body, &response); err != nil {
			log.Printf("Error unmarshaling response: %v", err)
			return
		}

		if response.Status == "ok" {
			log.Printf("Successfully collected %d new repositories", len(response.Added))
		}
	})

	s.StartAsync()
}
