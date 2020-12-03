package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan string)
var ch2 = make(chan string)
var done = make(chan struct{})
var count uint64

func init() {
	go func() {
	loop1:
		for {
			select {
			case msg := <-ch2:
				count++
				ch1 <- msg
			case <-done:
				break loop1
			}
		}
	}()
	go func() {
	loop2:
		for {
			select {
			case ch2 <- <-ch1:
			case <-done:
				break loop2
			}
		}
	}()
}

func main() {
	fmt.Println("[INFO] begin")
	tick := time.Tick(time.Second * 1)
	ch1 <- "message"
	<-tick
	close(done)
	fmt.Println("[INFO] end")
	fmt.Printf("[INFO] ping-pong for %d times\n", count)
}
