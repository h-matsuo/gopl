package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/h-matsuo/gopl/ch02/ex05/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	count := 0
	for i := 0; i < 8; i++ {
		count += popcount.BitClearPopCount(uint64(c1[i] ^ c2[i]))
	}

	fmt.Printf("c1: %x\nc2: %x\n# of different bits: %d\n", c1, c2, count)
}
