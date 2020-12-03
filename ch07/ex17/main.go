package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				names := []string{}
				for _, s := range stack {
					names = append(names, s.Name.Local)
				}
				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if strings.HasPrefix(y[0], "#") {
			var idName string
			for _, attr := range x[0].Attr {
				if attr.Name.Local != "id" {
					continue
				}
				idName = attr.Value
				break
			}
			if idName == y[0][1:] {
				y = y[1:]
			}
		} else if strings.HasPrefix(y[0], ".") {
			var className string
			for _, attr := range x[0].Attr {
				if attr.Name.Local != "class" {
					continue
				}
				className = attr.Value
				break
			}
			if className == y[0][1:] {
				y = y[1:]
			}
		} else {
			if x[0].Name.Local == y[0] {
				y = y[1:]
			}
		}
		x = x[1:]
	}
	return false
}
