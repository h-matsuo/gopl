package main

import (
	"fmt"
	"os"

	"github.com/h-matsuo/gopl/ch07/ex14/eval"
)

func main() {
	expr, err := eval.Parse("10 * { pow(x,2), 1+2, + } / 5") // 10 * ( pow(b,2) + pow(b,2) + pow(b,2) ) / 5
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(expr.String())
	fmt.Printf("==> %f\n", expr.Eval(eval.Env{"x": 3}))
}
