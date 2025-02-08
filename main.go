package main

import (
	"encoding/json"
	"net/http"

	"causeurgnocchi/dictionary/scraper"
)

func main() {
	http.HandleFunc("/api/vocabularies/{search}", getVocabulary)
	http.ListenAndServe(":8080", nil)
}

func getVocabulary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	json.NewEncoder(w).Encode(scraper.GetVocabulary(r.PathValue("search")))
}
