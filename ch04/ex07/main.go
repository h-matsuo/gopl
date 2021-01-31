package main

import (
	"fmt"
)

func main() {
	b := []byte("Hello, 世界")
	reverse(b)
	fmt.Println(string(b))
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
