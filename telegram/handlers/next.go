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

type repo struct {
	ID     int    `json:"id"`
	Posted bool   `json:"posted"`
	URL    string `json:"url"`
	Text   string `json:"text"`
}

type responseData struct {
	All      int    `json:"all"`
	Posted   int    `json:"posted"`
	Unposted int    `json:"unposted"`
	Items    []repo `json:"items"`
}

type apiResponse struct {
	Data    responseData `json:"data"`
	Message string       `json:"message"`
	Status  string       `json:"status"`
}

func NextPostHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	if userID != config.ADMIN_ID {
		return
	}

	limit := 1
	msgParts := strings.Fields(update.Message.Text)
	if len(msgParts) > 1 {
		if num, err := strconv.Atoi(msgParts[1]); err == nil && num > 0 {
			limit = num
		}
	}

	requestBody := fmt.Sprintf(`{"limit": %d, "posted": false}`, limit)
	req, err := http.NewRequest("POST", config.CHAPPIE_SERVER_URL+"/api/get-repository/", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		log.Panicln(err)
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.CHAPPIE_SERVER_BEARER)

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

	isDisabled := true
	for _, repo := range apiResponse.Data.Items {
		message := fmt.Sprintf("%s\n\n%s", repo.URL, repo.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:             userID,
			Text:               message,
			LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: &isDisabled},
		})
	}
}
