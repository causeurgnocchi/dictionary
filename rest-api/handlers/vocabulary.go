package handlers

import "encoding/json"

type vocabulary struct {
	Characters []character
	Meanings   []string
}

func (v vocabulary) Writing() string {
	writing := ""
	for _, c := range v.Characters {
		writing += c.Writing
	}
	return writing
}

func (v vocabulary) Reading() string {
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

func (v vocabulary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Characters []character `json:"characters"`
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
