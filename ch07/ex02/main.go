package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w, n := CountingWriter(os.Stdout)
	w.Write([]byte("Hello, world!\n")) // Also writes to os.Stdout
	fmt.Printf("written bytes: %d\n", *n)
}

type counter struct {
	w io.Writer
	n int64
}

func (c *counter) Write(p []byte) (int, error) {
	c.w.Write(p)
	c.n += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c counter
	c.w = w
	return &c, &c.n
}
