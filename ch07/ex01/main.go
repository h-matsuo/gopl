package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	var c WordLineCounter
	c.Write([]byte(`Hello,
	this is a sample.
	Byebyte.`))
	fmt.Printf("words: %d, lines: %d\n", c.NumWords, c.NumLines)
}

type WordLineCounter struct {
	NumWords int
	NumLines int
}

func (c *WordLineCounter) Write(p []byte) (int, error) {
	// Count lines
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		c.NumLines++
	}
	// Count words
	scanner = bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		c.NumWords++
	}
	return len(p), nil
}
