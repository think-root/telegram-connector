package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"telegram-bridge/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func ManualAddPostHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		messageText := update.Message.Text
		repoURLs := strings.TrimPrefix(messageText, "/add")
		repoURLs = strings.TrimSpace(repoURLs)

		if repoURLs == "" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text:   "Please provide at least one GitHub repository URL",
			})
			return
		}

		requestBody := struct {
			URL string `json:"url"`
		}{
			URL: repoURLs,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			log.Println(err)
			return
		}

		req, err := http.NewRequest("POST", config.CHAPPIE_SERVER_URL+"think-root/api/manual-generate/", bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Add("Authorization", "Bearer "+config.CHAPPIE_SERVER_BEARER)
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		var response struct {
			Status    string   `json:"status"`
			Added     []string `json:"added"`
			DontAdded []string `json:"dont_added"`
		}

		if err := json.Unmarshal(body, &response); err != nil {
			log.Println(err)
			return
		}

		if response.Status == "ok" {
			var message strings.Builder
			if len(response.Added) > 0 {
				message.WriteString("✅ Added:\n" + strings.Join(response.Added, "\n") + "\n\n")
			}
			if len(response.DontAdded) > 0 {
				message.WriteString("❌ Not Added:\n" + strings.Join(response.DontAdded, "\n"))
			}
			if message.Len() > 0 {
				isDisabled := true
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:             userID,
					Text:               message.String(),
					LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: &isDisabled},
				})
			}
		}
	}
}
