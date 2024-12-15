package main

import (
	"causeurgnocchi/dictionary/handlers"
	"net/http"
)

func main() {
	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/", &handlers.BaseHandler{})
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.Handle("/results", http.StripPrefix("/results", &handlers.ResultsHandler{}))
	http.Handle("/api/vocabularies/{search}", http.StripPrefix("/api/vocabularies/", &handlers.VocabulariesApiHandler{}))
	http.ListenAndServe(":8080", nil)
}
