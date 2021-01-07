package main

import (
	"fmt"
	"os"

	"github.com/h-matsuo/gopl/ch10/ex02/loader"
	_ "github.com/h-matsuo/gopl/ch10/ex02/loader/tar" // Register decoder for tar format
	_ "github.com/h-matsuo/gopl/ch10/ex02/loader/zip" // Register decoder for zip format
)

func main() {
	if err := loader.Load(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
