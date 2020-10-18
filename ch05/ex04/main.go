package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type config struct {
	a, img, script, stylesheet bool
}

var (
	fA          = flag.Bool("a", false, "Collect <a> elements' `href` links")
	fImg        = flag.Bool("img", false, "Collect <img> elements' `src` links")
	fScript     = flag.Bool("script", false, "Collect <script> elements' `src` links")
	fStylesheet = flag.Bool("stylesheet", false, "Collect <link rel=\"stylesheet\"> elements' `href` links")
	flags       = &config{}
)

func main() {
	flag.Parse()
	flags = &config{*fA, *fImg, *fScript, *fStylesheet}
	if !flags.a && !flags.img && !flags.script && !flags.stylesheet {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], "Error: specify at least 1 option\n")
		os.Exit(1)
	}

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" && flags.a {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		if n.Data == "img" && flags.img {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		if n.Data == "script" && flags.script {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		if n.Data == "link" && flags.stylesheet {
			href := ""
			containsStylesheet := false
			for _, attr := range n.Attr {
				if attr.Key == "rel" && attr.Val == "stylesheet" {
					containsStylesheet = true
				}
				if attr.Key == "href" {
					href = attr.Val
				}
			}
			if containsStylesheet {
				links = append(links, href)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
