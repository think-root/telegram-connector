package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"telegram-connector/telegram/services"

	"github.com/go-telegram/bot"
)

// logRequests is a middleware that logs incoming HTTP request details
func logRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] Path: %s | User-Agent: %s | Content-Type: %s",
			r.Method,
			r.URL.Path,
			r.Header.Get("User-Agent"),
			r.Header.Get("Content-Type"),
		)
		next.ServeHTTP(w, r)
	}
}

// Start initializes and runs the HTTP server with telegram message endpoint
func Start(port string, ctx context.Context, b *bot.Bot) {
	http.HandleFunc("/telegram/send-message", logRequests(func(w http.ResponseWriter, r *http.Request) {
		// Allow only POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Validate API key from request headers
		tokenHeader := r.Header.Get("X-API-Key")
		tokenEnv := os.Getenv("X_API_KEY")
		if tokenHeader == "" || tokenEnv == "" || tokenHeader != tokenEnv {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Parse multipart form data with 10MB limit
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Extract and validate required form fields
		url := r.FormValue("url")
		text := r.FormValue("text")
		if url == "" || text == "" {
			http.Error(w, "Missing url or text", http.StatusBadRequest)
			return
		}

		// Get and process image file from form data
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		imageData, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading image file", http.StatusInternalServerError)
			return
		}

		// Send message to Telegram
		services.SendMessage(ctx, b, url, text, imageData)

		// Return success response
		resp := map[string]string{"status": "OK", "message": "Message sent successfully"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))

	log.Printf("server: running on port %s", port)
	http.ListenAndServe(":"+port, nil)
}
