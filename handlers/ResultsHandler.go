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
    Results []vocabulary
    LastSearch string
}

type vocabulary struct {
    Characters []character
    Meanings []string `selector:".concept_light-meanings.medium-9.columns .meaning-meaning:not(:has(*))"`
}

func (r vocabulary) Writing() string {
    w := ""
    for _, c := range(r.Characters) {
        w += c.Writing
    }
    return w
}

func (r vocabulary) Reading() string {
    reading := ""
    for _, c := range(r.Characters) {
        reading += c.Reading
    }
    return reading
}

type character struct {
	Writing string
	Reading string
}

func (h ResultsHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
    results := make([]vocabulary, 0)
    collector := colly.NewCollector()
	collector.OnHTML("#primary", func(container *colly.HTMLElement) {
        container.ForEach(".concept_light.clearfix", func(i int, element *colly.HTMLElement) {
			writing := element.ChildText(`.concept_light-representation > .text`)

			furigana := make([]string, 0)
			element.ForEach(`.concept_light-representation > .furigana > span`, func(i int, e *colly.HTMLElement) {
				furigana = append(furigana, e.Text)
			})
            
			characters := make([]character, 0)
			for i, w := range([]rune(writing)) {
                f := furigana[i]
                if f == "" {
                    f = string(w)
                }
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
            results = append(results, v)
        })
	})

    collector.Visit("https://jisho.org/search/" + req.URL.Query().Get("search"))
    
	funcMap := template.FuncMap{
		"attr": func(attr string) template.HTMLAttr { return template.HTMLAttr(attr) },
	}
    tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("assets/html/base.html", "assets/html/results.html"))
    tmpl.ExecuteTemplate(writer, "base", resultsPageData{
		Results: results,
		LastSearch: req.URL.Query().Get("search"),
	})
}