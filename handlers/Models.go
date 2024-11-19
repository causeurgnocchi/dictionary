package handlers

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