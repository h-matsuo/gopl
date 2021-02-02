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
	// fmt.Printf("%v\n", root)
	print(root)
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
			// fmt.Println("[DEBUG] StartElement: <" + tok.Name.Local + ">")
			parent := &Element{tok.Name, tok.Attr, nil}
			children := visit(dec)
			parent.Children = children
			node = append(node, parent)
		case xml.EndElement:
			// fmt.Println("[DEBUG] EndElement: </" + tok.Name.Local + ">")
			return node
		case xml.CharData:
			// fmt.Println("[DEBUG] CharData: " + CharData(tok))
			node = append(node, CharData(tok))
		}
	}
}

func print(node []Node) {
	for _, n := range node {
		fmt.Println("[DEBUG] node")
		switch e := n.(type) {
		case CharData:
			fmt.Print(e)
		case Element:
			fmt.Print("<" + e.Type.Local + " ")
			for _, a := range e.Attr {
				fmt.Print(a.Name.Local + "=" + a.Value + " ")
			}
			fmt.Print(">")
			print(e.Children)
			fmt.Print("</" + e.Type.Local + ">")
		}
	}
}
