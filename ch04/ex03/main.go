package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(ptr *[6]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}
