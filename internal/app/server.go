package app

import (
	"io"
	"log"
	"net/http"
)

var hashTable = make(map[string]string)

func ShortenerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		if r.URL.Path != "/" {
			log.Println("Bad Request: " + r.URL.Path)
			http.Error(w, "Bad Request. Only / allowed.", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error at io.ReadAll: " + err.Error())
			http.Error(w, "Error at io.ReadAll: "+err.Error(), http.StatusBadRequest)
			return
		}
		longURL := string(body)
		shortURL := GetHash(longURL)

		// индекс есть в таблице
		if _, ok := hashTable[shortURL]; ok {
			// а строки разные - катастрофа
			if longURL != hashTable[shortURL] {
				err := "Double hash " + shortURL + " for strings: <" + longURL + "> & <" + hashTable[shortURL] + ">"
				log.Println(err)
				http.Error(w, err, http.StatusInternalServerError)
				return
			}
		} else {
			hashTable[shortURL] = longURL
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortURL))

	case http.MethodGet:
		if r.URL.Path == "/" {
			http.Error(w, "Bar Request", http.StatusBadRequest)
			return
		}
		shortURL := r.URL.Path[1:]
		if longURL, ok := hashTable[shortURL]; ok {
			longURL = "/" + longURL
			w.Header().Set("Location", longURL)
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			http.Error(w, "Long URL for short <"+shortURL+"> not found", http.StatusNotFound)
		}

	default:
		http.Error(w, "Bar Request", http.StatusBadRequest)

	}
}
