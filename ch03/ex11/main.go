package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	// Process sign
	if s != "" && (s[0] == '+' || s[0] == '-') {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	// Process interger part
	integer := s
	idxPoint := strings.Index(s, ".")
	if idxPoint >= 0 {
		integer = s[:idxPoint]
	}
	extra := len(integer) % 3
	if extra != 0 {
		buf.WriteString(integer[:extra])
	}
	for i := extra; i < len(integer); i = i + 3 {
		if i != 0 && i+2 < len(integer) {
			buf.WriteRune(',')
		}
		buf.WriteString(integer[i : i+3])
	}
	// Process others
	buf.WriteString(s[idxPoint:])
	return buf.String()
}
