package service

import (
	"encoding/json"
	"net/http"

	"github.com/artem-xox/go-shorty/internal/store"
)

type LinkMessage struct {
	URL string `json:"url"`
}

func (s *ShortyService) SetLink(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var link LinkMessage
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storeLink := store.NewLink(link.URL)
	s.Logger.Info(storeLink)

	err = s.Store.SetLink(r.Context(), storeLink)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storeLink)
}

func (s *ShortyService) GetLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedURL := r.URL.Path[len("/l/"):]
	storeLink := store.Link{Hash: hashedURL}
	s.Logger.Info(storeLink)

	retrievedLink, err := s.Store.GetLink(r.Context(), storeLink)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, retrievedLink.Long, http.StatusFound)
}
