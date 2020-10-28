package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

var host string

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlStr string) []string {
	if u, err := url.Parse(urlStr); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Invalid url: %s\n", urlStr)
	} else if u.Hostname() == host {
		fmt.Printf("[DOWNLOAD] %s\n", urlStr)
		download(u)
	} else {
		fmt.Printf("[SKIPPED] %s\n", urlStr)
	}

	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
}

func download(u *url.URL) {
	filePath := u.Hostname() + u.Path
	if u.Path == "" {
		filePath += "/index.html"
	} else if strings.HasSuffix(filePath, "/") {
		filePath += "index.html"
	}

	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create dir: %s\n", path.Dir(filePath))
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create file: %s\n", filePath)
		return
	}
	defer file.Close()

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to download file: %s\n", filePath)
		return
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to save file: %s\n", filePath)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: Specify url\n", os.Args[0])
		os.Exit(1)
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	host = u.Hostname()
	breadthFirst(crawl, []string{os.Args[1]})
}
