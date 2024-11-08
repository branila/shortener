package main

import (
	"log"
	"net/http"

	"github.com/branila/shortener/db"
	"github.com/branila/shortener/handlers"
)

func main() {
	db.Init()

	http.HandleFunc("GET /", handlers.Redirect)
	http.HandleFunc("GET /shorten", handlers.ShortenPage)
	http.HandleFunc("POST /shorten", handlers.ShortenApi)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
