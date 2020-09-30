package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	extra := len(s) % 3
	if extra != 0 {
		buf.WriteString(s[:extra])
	}
	for i := extra; i < len(s); i = i + 3 {
		if i != 0 && i+2 < len(s) {
			buf.WriteRune(',')
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
