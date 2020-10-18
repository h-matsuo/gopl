package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode && len(n.Data) > 0 {
		texts = append(texts, n.Data)
	}
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		visit(texts, n.FirstChild)
	}
	visit(texts, n.NextSibling)
	return texts
}
