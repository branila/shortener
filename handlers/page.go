package handlers

import (
	"html/template"
	"net/http"
)

func ShortenPage(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	view.Execute(w, nil)
}
