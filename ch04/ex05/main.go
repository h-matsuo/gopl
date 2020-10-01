package main

import (
	"fmt"
)

func main() {
	fmt.Println(distinct([]string{}))
	fmt.Println(distinct([]string{"a"}))
	fmt.Println(distinct([]string{"a", "a"}))
	fmt.Println(distinct([]string{"a", "b", "b", "ab", "ab"}))
	fmt.Println(distinct([]string{"a", "b", "b", "a", "ab", "ab", "abc"}))
}

func distinct(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	prev := s[0]
	for i := 1; i < len(s); i++ {
		current := s[i]
		if current == prev {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
		prev = current
	}
	return s
}
