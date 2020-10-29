package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/h-matsuo/gopl/ch07/ex15/eval"
)

func main() {
	fmt.Printf("Enter expression: ")
	reader := bufio.NewReader(os.Stdin)
	exprStr, _ := reader.ReadString('\n')
	exprStr = strings.Replace(exprStr, "\n", "", -1)

	expr, err := eval.Parse(exprStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(expr.String())
	fmt.Printf("==> %f\n", expr.Eval(eval.Env{}))
}
