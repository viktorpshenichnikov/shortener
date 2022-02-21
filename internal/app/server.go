package main

import (
	"github.com/viktorpshenichnikov/shortener/internal/app"
	"io"
	"log"
	"net/http"
)

var hashTable = make(map[string]string)

func ShortenerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := r.URL.Path[1:]
		shortURL := app.GetHash(path)

		// индекс есть в таблице
		if _, ok := hashTable[shortURL]; ok {
			// а строки разные - катастрофа ((
			if path != hashTable[shortURL] {
				err := "Double hash " + shortURL + " for strings: <" + path + "> & <" + hashTable[shortURL] + ">"
				log.Println(err)
				http.Error(w, err, http.StatusInternalServerError)
				return
			}
		} else {
			hashTable[shortURL] = path
		}
		shortURL = "/" + shortURL
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Header().Set("Location", shortURL)

		w.Write([]byte("Path: " + path + " Location: " + shortURL))

	case http.MethodPost:
		longURL, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error at io.ReadAll: " + err.Error())
			http.Error(w, "Error at io.ReadAll: "+err.Error(), http.StatusBadRequest)
			return
		}
		shortURL
	default:
		http.Error(w, "Bar Request", http.StatusBadRequest)

	}
}
