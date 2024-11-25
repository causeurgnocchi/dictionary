package handlers

import (
	"encoding/json"
	"net/http"
)

type VocabulariesApiHandler struct {
}

func (h VocabulariesApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scrapeJisho(r.PathValue("search")));
}