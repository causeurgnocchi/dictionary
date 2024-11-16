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

type vocabulary struct {
    Characters []character
    Meanings []string `selector:".concept_light-meanings.medium-9.columns .meaning-meaning:not(:has(*))"`
}
func (v vocabulary) Writing() string {
    w := ""
    for _, c := range(v.Characters) {
        w += c.Writing
    }
    return w
}
func (v vocabulary) Reading() string {
    r := ""
    for _, c := range(v.Characters) {
        r += c.Reading
    }
    return r
}

type character struct {
	Writing string
	Reading string
}

func (h ResultsHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
    vocabularies := make([]vocabulary, 0)
    
    collector := colly.NewCollector()
	collector.OnHTML("#primary .concept_light.clearfix", func(element *colly.HTMLElement) {
        writing := element.ChildText(`.concept_light-representation > .text`)

        furigana := make([]string, 0)
        element.ForEach(`.concept_light-representation > .furigana > span`, func(i int, e *colly.HTMLElement) {
            furigana = append(furigana, e.Text)
        })
        
        characters := make([]character, 0)
        for i, w := range([]rune(writing)) {
            f := furigana[i]
            c := character{
                Writing: string(w),
                Reading: f, 
            }
            characters = append(characters, c)
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