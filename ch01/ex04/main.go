package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/h-matsuo/gopl/ch01/ex04/counter"
)

// Extracted for unit tests
var (
	in     io.Reader = os.Stdin
	out    io.Writer = os.Stdout
	osArgs []string  = os.Args
)

func main() {
	lc := counter.NewLineCounter()
	files := osArgs[1:]
	if len(files) == 0 {
		countLines(in, lc)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			func() {
				defer f.Close()
				countLines(f, lc)
			}()
		}
	}
	showResult(lc)
}

func countLines(r io.Reader, lc *counter.LineCounter) {
	input := bufio.NewScanner(r)
	fileName := os.Stdin.Name()
	if f, ok := r.(*os.File); ok {
		fileName = f.Name()
	}
	for input.Scan() {
		lc.AddLine(input.Text(), fileName)
	}
}

func showResult(lc *counter.LineCounter) {
	for _, c := range lc.ToSlice() {
		if c.Occurrences() > 1 {
			fmt.Fprintf(out, "%d\t%v\t%s\n", c.Occurrences(), c.FileNames(), c.Line())
		}
	}
}
