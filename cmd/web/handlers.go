package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"ozonTest/pkg/models"
)

// На каждую страницу сделать ограничение на методы запросов (get, post, put и тд)
func createShortLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var tmpLink models.Link
	if err := json.NewDecoder(r.Body).Decode(&tmpLink); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := url.ParseRequestURI(tmpLink.Long)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Введенная строка не является ссылкой."))
		return
	}

	link := models.NewLink(tmpLink.Long)
	link.Short, _ = Storage.CreateShortLink(link)
	json.NewEncoder(w).Encode(link)
}

func getLongLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(405)
	}

	var tmpLink models.Link
	if err := json.NewDecoder(r.Body).Decode(&tmpLink); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longLink, err := Storage.GetLongLink(tmpLink.Short)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	currentLink := models.Link{Long: longLink, Short: tmpLink.Short}
	json.NewEncoder(w).Encode(&currentLink)
}
