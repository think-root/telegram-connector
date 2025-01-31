package whatsapp

import (
	"bytes"
	"chappie_bot/config"
	"encoding/json"
	"net/http"
)

type WhatsAppRequest struct {
	Text string `json:"text"`
	Jid  string `json:"jid"`
	Type string `json:"type"`
}

type WhatsAppResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SendMessageToWhatsApp(text string, jid string) bool {
	payload := WhatsAppRequest{
		Text: text,
		Jid:  jid,
		Type: "channel",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return false
	}

	req, err := http.NewRequest("POST", config.WAPP_SERVER_URL+"wapp/send-message", bytes.NewBuffer(jsonData))
	if err != nil {
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.WAPP_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var response WhatsAppResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false
	}

	return response.Success
}