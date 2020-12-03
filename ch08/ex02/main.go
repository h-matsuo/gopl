package main

import (
	"flag"

	"github.com/h-matsuo/gopl/ch08/ex02/pi"
)

var port = flag.Int("port", 21, "port number")

func main() {
	flag.Parse()
	pi.StartProtocolInterpreter(*port)
}
