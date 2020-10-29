package main

import (
	"fmt"
	"os"

	"github.com/h-matsuo/gopl/ch07/ex13/eval"
)

func main() {
	expr, err := eval.Parse("sqrt(A/pi) * (pow(x, sqrt(3))+pow(sqrt(sqrt(y)), 3)) * (F-(-32)) / (5*-9)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.String())

	_, err = eval.Parse(expr.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("OK")
	}
}
