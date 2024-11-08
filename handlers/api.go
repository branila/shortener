package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/branila/shortener/db"
)

func generateUniqueId() string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	for {
		id := ""
		for range 6 {
			id += string(chars[rand.Intn(len(chars))])
		}

		_, exists := db.GetId(id)
		if !exists {
			return id
		}
	}
}

type ShortenRequest struct {
	Url string
}

func ShortenApi(w http.ResponseWriter, r *http.Request) {
	var data ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	url := data.Url

	if url == "" {
		http.Error(w, "Bad request: URL is required", http.StatusBadRequest)
		return
	}

	re := regexp.MustCompile(`^https?://`)

	if !re.MatchString(url) {
		url = "http://" + url
	}

	response := map[string]string{
		"id": generateUniqueId(),
	}

	db.SetId(response["id"], url)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
