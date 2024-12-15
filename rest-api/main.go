package main

import (
	"causeurgnocchi/dictionary/handlers"
	"net/http"
)

func main() {
	http.Handle("/api/vocabularies/{search}", http.StripPrefix("/api/vocabularies/", &handlers.VocabulariesHandler{}))
	http.ListenAndServe(":8080", nil)
}
