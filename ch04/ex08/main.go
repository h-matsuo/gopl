package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			counts["invalid"]++
			continue
		}
		if unicode.IsControl(r) {
			counts["control"]++
		}
		if unicode.IsDigit(r) {
			counts["digit"]++
		}
		if unicode.IsGraphic(r) {
			counts["graphic"]++
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		}
		if unicode.IsLower(r) {
			counts["lower"]++
		}
		if unicode.IsMark(r) {
			counts["mark"]++
		}
		if unicode.IsNumber(r) {
			counts["number"]++
		}
		if unicode.IsPrint(r) {
			counts["print"]++
		}
		if unicode.IsPunct(r) {
			counts["punctuation"]++
		}
		if unicode.IsSpace(r) {
			counts["space"]++
		}
		if unicode.IsTitle(r) {
			counts["title"]++
		}
		if unicode.IsUpper(r) {
			counts["upper"]++
		}
	}
	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t\t%d\n", c, n)
	}
}
