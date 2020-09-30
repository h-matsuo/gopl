package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <STRING 1> <STRING 2>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(judgeAnagram(os.Args[1], os.Args[2]))
}

func judgeAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charSet1, charSet2 := map[rune]int{}, map[rune]int{}
	for _, r := range s1 {
		charSet1[r]++
	}
	for _, r := range s2 {
		charSet2[r]++
	}

	for r, n := range charSet1 {
		if n != charSet2[r] {
			return false
		}
	}
	return true
}
