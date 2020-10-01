package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

// rotate rotates `s` left by `n` positions.
func rotate(s []int, n int) {
	n = n % len(s)
	tmp := make([]int, n)
	copy(tmp, s[:n])
	copy(s[:len(s)-n], s[n:])
	copy(s[len(s)-n:], tmp)
}
