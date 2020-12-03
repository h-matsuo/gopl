package main

import (
	"fmt"
	"time"
)

const length = 10000000

var chs [length]chan string

func init() {
	fmt.Printf("[INFO] preparing for %d goroutines...\n", length)
	for i := 0; i < length; i++ {
		chs[i] = make(chan string)
	}
	for i := 0; i < length-1; i++ {
		go func(from <-chan string, to chan<- string) {
			to <- <-from
		}(chs[i], chs[i+1])
	}
}

func main() {
	fmt.Println("[INFO] begin")
	begin := time.Now()
	chs[0] <- "message"
	fmt.Println(<-chs[length-1])
	end := time.Now()
	fmt.Println("[INFO] end")
	fmt.Printf("[INFO] it takes %d msecs\n", end.Sub(begin).Milliseconds())
}
