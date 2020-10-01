package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	s384 = flag.Bool("sha384", false, "calculate SHA384")
	s512 = flag.Bool("sha512", false, "calculate SHA512")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if *s384 && *s512 || len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [ -sha384 | -sha512 ] <INPUT>\n", os.Args[0])
		os.Exit(1)
	}
	switch {
	case *s384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(args[0])))
	case *s512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(args[0])))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(args[0])))
	}
}
