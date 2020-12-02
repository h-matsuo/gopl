package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string, done <-chan struct{}) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

type response struct {
	url      string
	filename string
	n        int64
	err      error
}

func main() {
	responses := make(chan *response, len(os.Args[1:]))
	done := make(chan struct{})

	for _, url := range os.Args[1:] {
		go func(url string) {
			local, n, err := fetch(url, done)
			responses <- &response{url, local, n, err}
		}(url)
	}

	res := <-responses // get quickest response
	go func() {
		done <- struct{}{}
	}()

	if res.err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", res.url, res.err)
	}
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", res.url, res.filename, res.n)
}
