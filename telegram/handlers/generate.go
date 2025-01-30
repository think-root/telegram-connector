package handlers

import (
	"bytes"
	"chappie_bot/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

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

func GeneratePostsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		maxRepos := 1
		msgText := update.Message.Text
		parts := strings.Fields(msgText)

		if len(parts) > 1 {
			if num, err := strconv.Atoi(parts[1]); err == nil && num > 0 {
				maxRepos = num
			}
		}

		payload := generateRequest{
			MaxRepos:           maxRepos,
			Since:              "daily",
			SpokenLanguageCode: "en",
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Error marshaling request: %v", err)
			return
		}

		req, err := http.NewRequest("POST", config.CHAPPIE_SERVER_URL+"/api/auto-generate/", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+config.CHAPPIE_SERVER_BEARER)

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
			var message strings.Builder

			if len(response.Added) > 0 {
				message.WriteString("✅ Added repositories:\n")
				for _, repo := range response.Added {
					message.WriteString(fmt.Sprintf("• %s\n", repo))
				}
			}

			if len(response.DontAdded) > 0 {
				if message.Len() > 0 {
					message.WriteString("\n")
				}
				message.WriteString("❌ Not added repositories:\n")
				for _, repo := range response.DontAdded {
					message.WriteString(fmt.Sprintf("• %s\n", repo))
				}
			}

			if message.Len() > 0 {
				isDisabled := true
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: userID,
					Text:   message.String(),
					LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: &isDisabled},
				})
			}
		}
	}
}
