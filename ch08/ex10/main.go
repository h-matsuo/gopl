package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string, canceled <-chan struct{}) []string {
	// fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := Extract(url, canceled)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Handle SIGINT
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	canceled := make(chan struct{})
	go func() {
		<-sigs
		fmt.Println("canceling http requests...")
		canceled <- struct{}{}
		// os.Exit(0)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
loop:
	for ; n > 0; n-- {
		list := []string{}
		select {
		case list = <-worklist:
		case <-canceled:
			break loop
		}
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link, canceled)
				}(link)
			}
		}
	}
}
