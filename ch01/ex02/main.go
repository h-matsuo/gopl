package main

import (
	"fmt"
	"io"
	"os"
)

// Extracted for unit tests
func printStrings(w io.Writer, strs []string) {
	for idx, arg := range strs[1:] {
		fmt.Fprintf(w, "(%d, %s)\n", idx, arg)
	}
}

func main() {
	printStrings(os.Stdout, os.Args)
}
