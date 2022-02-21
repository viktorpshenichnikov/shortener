package main

import (
	"github.com/viktorpshenichnikov/shortener/internal/app"
	"net/http"
)

func main() {
	http.HandleFunc("/", app.ShortenerHandler)
	http.ListenAndServe("localhost:8080", nil)
}
