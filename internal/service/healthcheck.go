package service

import "net/http"

func (s *ShortyService) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
