package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Extracted for unit tests
func printStrings(w io.Writer, strs []string) {
	fmt.Fprintf(w, strings.Join(strs, " "))
}

func main() {
	printStrings(os.Stdout, os.Args)
}
