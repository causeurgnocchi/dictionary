package scraper

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"causeurgnocchi/dictionary/models"
)

func GetVocabulary(searchTerm string) []models.Vocabulary {
	url := "https://jisho.org/search/" + searchTerm

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]models.Vocabulary, 0)
	vocabularyElements := filter(*doc, []Selector{
		{Type: idSelector, Val: "primary"},
		{Type: classSelector, Val: "concept_light clearfix"},
	})

	for _, v := range vocabularyElements {
		writing := []rune(childText(v, []Selector{
			{Type: classSelector, Val: "concept_light-representation"},
			{Type: classSelector, Val: "text"},
		}))

		furigana := make([]string, len(writing))
		furiganaElements := filter(v, []Selector{
			{Type: classSelector, Val: "concept_light-representation"},
			{Type: classSelector, Val: "furigana"},
			{Type: tagSelector, Val: "span"},
		})

		for i, f := range furiganaElements {
			if f.FirstChild != nil {
				furigana[i] = strings.TrimSpace(f.FirstChild.Data)
			}
		}

		characters := make([]models.Character, len(writing))
		for i, w := range writing {
			c := models.Character{
				Writing:  string(w),
				Furigana: furigana[i],
			}
			characters[i] = c
		}

		meanings := make([]string, 0)
		meaningsElements := filter(v, []Selector{
			{Type: classSelector, Val: "concept_light-meanings medium-9 columns"},
			{Type: classSelector, Val: "meaning-meaning"},
		})
		for _, m := range meaningsElements {
			if m.FirstChild.Type != html.TextNode {
				continue
			}
			meanings = append(meanings, strings.TrimSpace(m.FirstChild.Data))
		}

		result = append(result, models.Vocabulary{
			Characters: characters,
			Meanings:   meanings,
		})
	}
	return result
}
