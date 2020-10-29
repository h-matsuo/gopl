package main

import (
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	r := LimitReader(strings.NewReader("Hello, world!"), 5)
	io.Copy(os.Stdout, r)
}

type limitReader struct {
	r io.Reader
	n int64 // num bytes to be read
	i int64 // current position
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n, 0}
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	if r.i >= r.n {
		return 0, io.EOF
	}
	toBeRead := int64(math.Min(float64(len(p)), float64(r.n-r.i)))
	n, err = r.r.Read(p[:toBeRead])
	r.i += toBeRead
	return
}
