package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Extracted for unit tests
var (
	out    io.Writer = os.Stdout
	osArgs []string  = os.Args
)

func main() {
	fmt.Fprintf(out, strings.Join(osArgs, " "))
}
