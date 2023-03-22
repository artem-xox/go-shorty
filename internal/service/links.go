package service

import "net/http"

func (s *ShortyService) GetLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *ShortyService) SetLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
