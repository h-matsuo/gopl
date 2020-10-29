package main

import (
	"io"
	"os"
)

func main() {
	r := NewReader("Hello, world!")
	io.Copy(os.Stdout, r)
}

type stringReader struct {
	buf []byte
	i   int // current position
}

func NewReader(s string) io.Reader {
	return &stringReader{[]byte(s), 0}
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.buf) {
		return 0, io.EOF
	}
	n = copy(p, r.buf[r.i:])
	r.i += n
	return
}
