package main

import (
	"causeurgnocchi/dictionary/scraper"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/vocabularies/{search}", searchVocabularies)
	http.HandleFunc("/api/vocabularies/{search}", vocabulariesAPI)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("./assets/html/base.html"))
	tmpl = template.Must(tmpl.Parse(`{{define "content"}}{{end}}`))
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		panic(err)
	}
}

func searchVocabularies(w http.ResponseWriter, r *http.Request) {
	var (
		funcs = template.FuncMap{
			"furiganasize": func(furigana string) template.HTMLAttr {
				count := len([]rune(furigana))
				var size float32
				if count == 1 {
					size = 0.5
				} else {
					size = 1 / float32(count)
				}
				attr := fmt.Sprintf(`style="font-size:%.2fem"`, size)
				return template.HTMLAttr(attr)
			},
		}
		files = []string{
			"./assets/html/base.html",
			"./assets/html/vocabularies.html",
		}
		term = r.PathValue("search")
		vocabularies = scraper.Search(term)
	)
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "base", vocabularies)
	if err != nil {
		panic(err)
	}
}

func vocabulariesAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// uncomment line below if using GitHub Codespaces
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// uncomment line below if in a local environment
	// w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
	term := r.PathValue("search")
	results := scraper.Search(term)
	json.NewEncoder(w).Encode(results)
}
