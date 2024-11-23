package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

type VocabulariesApiHandler struct {
}

func (h VocabulariesApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
    collector.Visit("https://jisho.org/search/" + r.PathValue("search"))
    
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scrapeJisho(r.PathValue("search")));
}