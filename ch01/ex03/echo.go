package echo

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Extracted for benchmark
var (
	out    io.Writer = os.Stdout
	osArgs []string  = os.Args
)

func nonEfficientEcho() {
	s, sep := "", ""
	for _, arg := range osArgs {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintf(out, s)
}

func efficientEcho() {
	fmt.Fprintf(out, strings.Join(osArgs, " "))
}
