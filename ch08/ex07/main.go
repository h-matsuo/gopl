package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

var host string

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(urlStr string) []string {
	tokens <- struct{}{} // acquire a token
	defer func() {
		<-tokens // release the token
	}()

	if u, err := url.Parse(urlStr); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Invalid url: %s\n", urlStr)
	} else if u.Hostname() == host {
		fmt.Printf("[DOWNLOAD] %s\n", urlStr)
		download(u)
	} else {
		fmt.Printf("[SKIPPED] %s\n", urlStr)
	}

	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
}

var depth = flag.Int("depth", 1, "depth limitation for crawling")

type worklistDto struct {
	links []string
	depth int
}

func main() {
	flag.Parse()

	worklist := make(chan *worklistDto)
	var n int // number of pending sends to worklist

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "%s: Specify url\n", os.Args[0])
		os.Exit(1)
	}
	targetURL := flag.Args()[0]

	u, err := url.Parse(targetURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	host = u.Hostname()

	// Start with the command-line arguments.
	n++
	go func() { worklist <- &worklistDto{[]string{targetURL}, 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > *depth {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- &worklistDto{crawl(link), depth + 1}
				}(link, list.depth)
			}
		}
	}
}
