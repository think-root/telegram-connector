package x

import (
	"bytes"
	"chappie_bot/config"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func CreateXPost(text, repoUrl, filename string) bool {
	apiKey := config.X_API_KEY
	if apiKey == "" {
		return false
	}

	url := config.X_URL + "/x/api/posts/create"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("text", text); err != nil {
		return false
	}

	if err := writer.WriteField("url", repoUrl); err != nil {
		return false
	}

	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	part, err := writer.CreateFormFile("image", filepath.Base(filename))
	if err != nil {
		return false
	}

	if _, err = io.Copy(part, file); err != nil {
		return false
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return false
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
