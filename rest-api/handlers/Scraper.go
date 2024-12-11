package handlers

import (
	"strings"

	"golang.org/x/net/html"
)

type selectorType int

const (
	classSelector selectorType = iota
	idSelector
	tagSelector
)

type selector struct {
	Type selectorType
	Val  string
}

func childText(n html.Node, selectors []selector) string {
	var text []rune
	for _, child := range filter(n, selectors) {
		for d := range child.Descendants() {
			if d.Type == html.TextNode {
				text = append(text, []rune(strings.TrimSpace(d.Data))...)
			}
		}
	}
	return string(text)
}

func filter(n html.Node, selectors []selector) []html.Node {
	if match(n, selectors[0]) {
		if len(selectors) > 1 {
			return filterChildren(n, selectors[1:])
		}
		return append(filterChildren(n, selectors), n)
	}

	if n.FirstChild == nil {
		return make([]html.Node, 0)
	}

	return filterChildren(n, selectors)
}

func filterChildren(n html.Node, childSelectors []selector) []html.Node {
	var matches []html.Node
	for c := range n.ChildNodes() {
		matches = append(matches, filter(*c, childSelectors)...)
	}
	return matches
}

func match(n html.Node, s selector) bool {
	if s.Type == tagSelector && n.Type == html.ElementNode && n.Data == s.Val {
		return true
	}

	for _, a := range n.Attr {
		if a.Val == s.Val {
			if s.Type == classSelector && a.Key == "class" || s.Type == idSelector && a.Key == "id" {
				return true
			}
		}
	}
	return false
}
