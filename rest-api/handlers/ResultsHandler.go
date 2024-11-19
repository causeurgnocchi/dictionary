package handlers

import (
	"html/template"
	"net/http"

	"github.com/gocolly/colly"
)

type ResultsHandler struct {
	BaseHandler
}
type resultsPageData struct {
    Vocabularies []vocabulary
    LastSearch string
}

func (h ResultsHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
    vocabularies := make([]vocabulary, 0)
    
    collector := colly.NewCollector()
	collector.OnHTML("#primary .concept_light.clearfix", func(element *colly.HTMLElement) {
        writing := []rune(element.ChildText(`.concept_light-representation > .text`))

        furigana := make([]string, len(writing))
        element.ForEach(`.concept_light-representation > .furigana > span`, func(i int, e *colly.HTMLElement) {
            furigana = append(furigana, e.Text)
        })
        
        characters := make([]character, len(writing))
		
        for i, w := range(writing) {
            f := furigana[i]
            c := character{
                Writing: string(w),
                Reading: f, 
            }
            characters[i] = c
        }

        v := vocabulary{
            Characters: characters,
        }
        element.Unmarshal(&v)
        vocabularies = append(vocabularies, v)
	})
    collector.Visit("https://jisho.org/search/" + req.URL.Query().Get("search"))
    
	funcMap := template.FuncMap{
		"attr": func(attr string) template.HTMLAttr { return template.HTMLAttr(attr) },
		"strlen": func(s string) int { return len([]rune(s)) },
	}
    tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("assets/html/base.html", "assets/html/results.html"))
    tmpl.ExecuteTemplate(writer, "base", resultsPageData{
		Vocabularies: vocabularies,
		LastSearch: req.URL.Query().Get("search"),
	})
}