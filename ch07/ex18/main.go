package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	root := visit(dec)
	fmt.Printf("%v\n", root)
}

func visit(dec *xml.Decoder) []Node {
	node := []Node{}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return node
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if node != nil {
				node = append(node, &Element{tok.Name, tok.Attr, []Node{}})
			}
			child := visit(dec)
			node = append(node, child)
		case xml.EndElement:
			return []Node{node}
		case xml.CharData:
			node = append(node, CharData(tok))
		}
	}
}
