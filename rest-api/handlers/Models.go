package handlers

import "encoding/json"

type vocabulary struct {
    Characters []character
    Meanings []string
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
		if (c.Furigana != "") {
			r += c.Furigana;
		} else {
			r += c.Writing;
		}
    }
    return r
}

func (v vocabulary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Characters []character `json:"characters"`
		Meanings []string `json:"meanings"`
		Writing string `json:"writing"`
		Reading string `json:"reading"`
	}{
		Characters: v.Characters,
		Meanings: v.Meanings,
		Writing: v.Writing(),
		Reading: v.Reading(),
	})
}

type character struct {
	Writing string `json:"writing"`
	Furigana string `json:"furigana"`
}