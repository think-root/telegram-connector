package x

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"telegram-connector/config"
)

func CreateXPost(text, repoUrl, filename string) bool {
	apiKey := config.X_API_KEY
	if apiKey == "" {
		log.Println("X API key is empty")
		return false
	}

	url := config.X_URL + "/x/api/posts/create"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("text", text); err != nil {
		log.Printf("Failed to write text field: %v", err)
		return false
	}

	if err := writer.WriteField("url", repoUrl); err != nil {
		log.Printf("Failed to write url field: %v", err)
		return false
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Failed to open file %s: %v", filename, err)
		return false
	}
	defer file.Close()

	part, err := writer.CreateFormFile("image", filepath.Base(filename))
	if err != nil {
		log.Printf("Failed to create form file: %v", err)
		return false
	}

	if _, err = io.Copy(part, file); err != nil {
		log.Printf("Failed to copy file contents: %v", err)
		return false
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return false
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d", resp.StatusCode)
		return false
	}

	return true
}
