package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

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

	// Replace a path with hostname to local path
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to read file: %s\n", filePath)
		return
	}
	hostnameReplacedHTMLText := replaceHostname(string(buf))

	if _, err = io.Copy(file, strings.NewReader(hostnameReplacedHTMLText)); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to save file: %s\n", filePath)
		return
	}
}
