package handlers

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func scrapeJisho(search string) []vocabulary {
    url := "https://jisho.org/search/" + search

    resp, err := http.Get(url)
    if (err != nil) {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if (err != nil) {
        log.Fatal(err)
    }

    vocabularies := make([]vocabulary, 0)
    vocabularyElements := filter(doc, []selector{
        {Type: idSelector, Val: "primary"},
        {Type: classSelector, Val: "concept_light clearfix"},
    })

    for _, v := range(vocabularyElements) {
        writing := []rune(childText(v, []selector{
            {Type: classSelector, Val: "concept_light-representation"},
            {Type: classSelector, Val: "text"},
        }))

        furigana := make([]string, len(writing))
        furiganaElements := filter(v, []selector{
            {Type: classSelector, Val: "concept_light-representation"},
            {Type: classSelector, Val: "furigana"},
            {Type: tagSelector, Val: "span"},
        })

        for i, f := range furiganaElements {
            if (f.FirstChild != nil) {
				furigana[i] = strings.TrimSpace(f.FirstChild.Data)
			}
        }
        
        characters := make([]character, len(writing))
        for i, w := range writing {
            c := character{
                Writing: string(w),
                Reading: furigana[i], 
            }
            characters[i] = c
        }

        meanings := make([]string, 0)
        meaningsElements := filter(v, []selector{
            {Type: classSelector, Val: "concept_light-meanings medium-9 columns"},
            {Type: classSelector, Val: "meaning-meaning"},
        })
        for _, m := range meaningsElements {
            if (m.FirstChild.Type != html.TextNode) {
                continue
            }
            meanings = append(meanings, strings.TrimSpace(m.FirstChild.Data))
        }

        vocabularies = append(vocabularies, vocabulary{
            Characters: characters,
            Meanings: meanings,
        })
    }
	return vocabularies
}