package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

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

	// Start with the command-line arguments.
	n++
	go func() { worklist <- &worklistDto{flag.Args(), 0} }()

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
