package main

import (
	"bytes"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func replaceHostname(htmlText string) string {
	doc, err := html.Parse(strings.NewReader(htmlText))
	if err != nil {
		return htmlText
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" || !strings.Contains(a.Val, host) {
					continue
				}
				u, err := url.Parse(a.Val)
				if err != nil || !strings.Contains(u.Scheme, "http") || len(u.Port()) > 0 {
					continue
				}
				u.Scheme = ""
				u.Host = ""
				a.Val = u.String()
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	buf := new(bytes.Buffer)
	html.Render(buf, doc)
	return buf.String()
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
