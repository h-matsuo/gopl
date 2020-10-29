package main

import "fmt"

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/h-matsuo/gopl/ch07/ex06/tempconv"
// )

// var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

// func main() {
// 	flag.Parse()
// 	fmt.Println(*temp)
// }

type MyType struct{}

func (m MyType) String() string {
	return "hoge"
}
func main() {
	var m MyType
	fmt.Printf("%q\n", m)
}
