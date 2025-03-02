package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"telegram-bridge/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DBStatisticHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID == config.ADMIN_ID {
		requestBody := `{"limit": 1, "posted": false}`
		req, err := http.NewRequest("POST", config.CONTENT_ALCHEMIST_URL+"think-root/api/get-repository/", bytes.NewBuffer([]byte(requestBody)))
		if err != nil {
			log.Panicln(err)
			return
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.CONTENT_ALCHEMIST_BEARER)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Panicln(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panicln(err)
			return
		}

		var apiResponse apiResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Panicln(err)
			return
		}

		message := fmt.Sprintf("📊 Statistics:\n\nAll posts: %d\nPosted: %d\nUnposted: %d",
			apiResponse.Data.All,
			apiResponse.Data.Posted,
			apiResponse.Data.Unposted)

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   message,
		})
	}
}
