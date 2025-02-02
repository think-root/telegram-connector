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

	resp, err := postJSON(config.WAPP_SERVER_URL+"wapp/send-message", config.WAPP_TOKEN, payload)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var response WhatsAppResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false
	}

	return response.Success
}

func postJSON(url, token string, payload interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	return client.Do(req)
}