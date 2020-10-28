package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Extracted for unit tests
var (
	out io.Writer = os.Stdout
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
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

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		fmt.Fprintf(out, "%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.TextNode:
		fmt.Fprintf(out, "%*s%s\n", depth*2, "", n.Data)
	case html.ElementNode:
		attrs := func() string {
			attrs := []string{}
			for _, attr := range n.Attr {
				attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
			}
			if len(attrs) > 0 {
				return " " + strings.Join(attrs, " ")
			}
			return ""
		}()
		fmt.Fprintf(out, "%*s<%s%s", depth*2, "", n.Data, attrs)
		if n.FirstChild == nil {
			fmt.Fprintf(out, " /")
		}
		fmt.Fprintf(out, ">\n")
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
