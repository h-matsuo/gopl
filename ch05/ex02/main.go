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
	counts := make(map[string]int)
	visit(counts, doc)
	for nodeType, count := range counts {
		fmt.Printf("<%s>: %d\n", nodeType, count)
	}
}

func visit(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	visit(counts, n.FirstChild)
	visit(counts, n.NextSibling)
}
