package handlers

import (
	"html/template"
	"net/http"
)

type ResultsHandler struct {
	BaseHandler
}

type resultsPageData struct {
    Vocabularies []vocabulary
    LastSearch string
}

func (h ResultsHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{
		"attr": func(attr string) template.HTMLAttr { return template.HTMLAttr(attr) },
		"strlen": func(s string) int { return len([]rune(s)) },
	}
    tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("assets/html/base.html", "assets/html/results.html"))
    tmpl.ExecuteTemplate(writer, "base", resultsPageData{
		Vocabularies: scrapeJisho(req.URL.Query().Get("search")),
		LastSearch: req.URL.Query().Get("search"),
	})
}