package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], "Error: specify input\n")
		os.Exit(1)
	}
	result := expand(os.Args[1], func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(result)
}

func expand(s string, f func(string) string) string {
	var scn scanner.Scanner
	scn.Init(strings.NewReader(s))
	// Treat leading '$' as part of an identifier
	// c.f. package scanner's example (IsIdentRune)
	scn.IsIdentRune = func(ch rune, i int) bool {
		return ch == '$' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	var previous, current int // buffer position
	var result strings.Builder

	for tok := scn.Scan(); tok != scanner.EOF; tok = scn.Scan() {
		if strings.HasPrefix(scn.TokenText(), "$") {
			current = scn.Position.Offset
			result.WriteString(s[previous:current])
			result.WriteString(f(scn.TokenText()[1:]))
			previous = current + len(scn.TokenText())
		}
	}
	result.WriteString(s[previous:])

	return result.String()
}
