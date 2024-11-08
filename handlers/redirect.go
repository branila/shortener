package handlers

import (
	"net/http"

	"github.com/branila/shortener/db"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/shorten", http.StatusFound)
		return
	}

	id := r.URL.Path[1:]

	url, exists := db.GetId(id)

	if !exists {
		http.Error(w, "URL Not Found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
