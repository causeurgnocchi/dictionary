package handlers

import (
	"bytes"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type selectorType int

const (
    classSelector selectorType = iota
    idSelector
    tagSelector
)

type selector struct{
    Type selectorType
    Val string
}

func childText(n *html.Node, s []selector) string {
	t := make([]rune, 0)
	children := filter(n, s)
	for _, c := range children {
        for d := range c.Descendants() {
            if (d.Type == html.TextNode) {
               t = append(t, []rune(strings.TrimSpace(d.Data))...)
            }
        }
    }
    return string(t)
}

func filter(n *html.Node, s []selector) []*html.Node {
    filterChildren := func(childSelectors []selector) []*html.Node {
        matches := make([]*html.Node, 0)
        for c := range n.ChildNodes() {
            matches = append(matches, filter(c, childSelectors)...)
        }
        return matches
    }
    
    if (match(n, s[0])) {
        if (len(s) > 1) {
            return filterChildren(s[1:])
        }
        return append(filterChildren(s), n)
    }
            
    if (n.FirstChild == nil) {
        return make([]*html.Node, 0)
    }

    return filterChildren(s)
}

func match(n *html.Node, s selector) bool {
    if (s.Type == tagSelector && n.Type == html.ElementNode && n.Data == s.Val) {
        return true
    }
    
    for _, a := range n.Attr {
        if (a.Val == s.Val) {
            if (s.Type == classSelector && a.Key == "class" ||  s.Type == idSelector && a.Key == "id") {
                return true
            }
        }
    }
    return false
}

func renderNode(n *html.Node) string {
    var b bytes.Buffer
    err := html.Render(&b, n)
    if (err != nil) {
        log.Fatal(err)
    }
    return b.String()
}