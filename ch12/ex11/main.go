package main

import (
	"fmt"
	"log"

	"github.com/h-matsuo/gopl/ch12/ex11/params"
)

func main() {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.Labels = []string{"foo", "bar"}
	data.MaxResults = 10
	packed, err := params.Pack(&data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf("/?%s\n", packed)
}
