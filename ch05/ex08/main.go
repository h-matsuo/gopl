package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], "Error: specify URL and id\n")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]

	if err := exec(url, id); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
		os.Exit(1)
	}
}

func exec(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	if found := ElementByID(doc, id); found != nil {
		fmt.Printf("ElementNode found: ")
		attrs := func() string {
			attrs := []string{}
			for _, attr := range found.Attr {
				attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
			}
			if len(attrs) > 0 {
				return " " + strings.Join(attrs, " ")
			}
			return ""
		}()
		fmt.Printf("<%s%s", found.Data, attrs)
		if found.FirstChild == nil {
			fmt.Printf(" /")
		}
		fmt.Printf(">\n")
	} else {
		fmt.Println("no such ElementNode found")
	}

	return nil
}

var targetID string
var foundNode *html.Node

func ElementByID(doc *html.Node, id string) *html.Node {
	targetID = id
	forEachNode(doc, startElement, nil)
	return foundNode
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(n) {
		return
	}
}

func startElement(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return true
	}
	for _, attr := range n.Attr {
		if attr.Key == "id" && attr.Val == targetID {
			foundNode = n
			return false
		}
	}
	return true
}
