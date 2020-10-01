package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/h-matsuo/gopl/ch04/ex13/omdb"
)

var (
	token = flag.String("token", "", "OMDb API key")
	title = flag.String("title", "", "Movie title to search")
	out   = flag.String("out", "", "File path to save the poster image")
)

func main() {

	flag.Parse()

	if *token == "" || *title == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "`-token`, `-title` and `-out` are required.")
		os.Exit(1)
	}

	res, err := omdb.SearchMovie(*token, *title)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if res.Error != "" {
		fmt.Fprintf(os.Stderr, "Error: %s\n", res.Error)
		os.Exit(1)
	}

	fmt.Printf("Download a poster of the movie \"%s\".\n", res.Title)
	fmt.Printf("Downloading from %s\n", res.Poster)
	download(res.Poster, *out)
}

func download(url, out string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "download failed: %s\n", resp.Status)
		os.Exit(1)
	}

	file, err := os.Create(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
}
