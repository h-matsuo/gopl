package main

import (
	"fmt"
	"io"
	"os"
)

// Extracted for unit tests
var (
	out    io.Writer = os.Stdout
	osArgs []string  = os.Args
)

func main() {
	for idx, arg := range osArgs[1:] {
		fmt.Fprintf(out, "(%d, %s)\n", idx, arg)
	}
}
