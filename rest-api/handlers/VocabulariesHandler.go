package handlers

import (
	"encoding/json"
	"net/http"
)

type VocabulariesHandler struct {
}

func (h VocabulariesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	json.NewEncoder(w).Encode(scrapeJisho(r.PathValue("search")))
}
