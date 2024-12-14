package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// urlMap holds the mapping of short URLs to long URLs
var urlMap = make(map[string]string)

func main() {
	http.HandleFunc("/", redirectURLHandler)
	http.HandleFunc("/shorten", shortenURLHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateShortCode() string {
	const shortCodeLen = 12
	randomBytes := make([]byte, shortCodeLen)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Unable to generate random bytes: " + err.Error())
	}

	shortCode := base64.RawURLEncoding.EncodeToString(randomBytes)

	if len(shortCode) > shortCodeLen {
		return shortCode[:shortCodeLen]
	}
	return shortCode
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()
	urlMap[shortCode] = longURL

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
	fmt.Fprintf(w, "Shortened URL: %s\n", shortenedURL)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {

	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	longURL, ok := urlMap[shortCode]
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
