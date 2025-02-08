package models

import "encoding/json"

type Vocabulary struct {
	Characters []Character
	Meanings   []string
}

func (v Vocabulary) Writing() string {
	writing := ""
	for _, c := range v.Characters {
		writing += c.Writing
	}
	return writing
}

func (v Vocabulary) Reading() string {
	reading := ""
	for _, c := range v.Characters {
		if c.Furigana != "" {
			reading += c.Furigana
		} else {
			reading += c.Writing
		}
	}
	return reading
}

func (v Vocabulary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Characters []Character `json:"characters"`
		Meanings   []string    `json:"meanings"`
		Writing    string      `json:"writing"`
		Reading    string      `json:"reading"`
	}{
		Characters: v.Characters,
		Meanings:   v.Meanings,
		Writing:    v.Writing(),
		Reading:    v.Reading(),
	})
}
