package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	ch := make(chan string)
	tick := time.Tick(time.Second * 10)
	done := make(chan struct{})

	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
		close(done)
	}()

loop:
	for {
		select {
		case text := <-ch:
			tick = time.Tick(time.Second * 10)
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, text, 1*time.Second)
			}()
		case <-tick:
			fmt.Fprintln(c, "No messages sent for 10 secs; timeout.")
			c.(*net.TCPConn).CloseRead()
			break loop
		case <-done:
			break loop
		}
	}

	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
