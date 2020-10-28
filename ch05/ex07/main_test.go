package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	tests := []struct {
		html string
	}{
		{`
		<html>
			<head>
				<link rel="stylesheet" href="foo.css">
			</head>
			<body>
				<!-- this is a comment -->
				<a href="foo.com">text</a>
				<div>
					<p><img src="foo.img"></p>
				</div>
			</body>
		</html>
		`},
	}

	for _, test := range tests {
		out = new(bytes.Buffer) // To capture output

		doc, _ := html.Parse(strings.NewReader(test.html))
		forEachNode(doc, startElement, endElement)
		got := out.(*bytes.Buffer).String()

		_, gotErr := html.Parse(strings.NewReader(got))

		if gotErr != nil {
			t.Errorf("forEachNode() prints invalid HTML, got error %v", gotErr)
		}
	}
}

func TestStartElement(t *testing.T) {
	tests := []struct {
		depth int
		node  *html.Node
		want  string
	}{
		{0, &html.Node{
			FirstChild: &html.Node{},
			Type:       html.ElementNode,
			Data:       "html",
			Attr:       nil,
		}, "<html>\n"},
		{1, &html.Node{
			FirstChild: nil,
			Type:       html.ElementNode,
			Data:       "img",
			Attr:       []html.Attribute{{Key: "src", Val: "foo.img"}},
		}, "  <img src=\"foo.img\" />\n"},
		{2, &html.Node{
			FirstChild: &html.Node{},
			Type:       html.ElementNode,
			Data:       "a",
			Attr:       []html.Attribute{{Key: "href", Val: "foo.com"}, {Key: "target", Val: "_blank"}},
		}, "    <a href=\"foo.com\" target=\"_blank\">\n"},
		{2, &html.Node{
			FirstChild: nil,
			Type:       html.CommentNode,
			Data:       "this is a comment",
			Attr:       nil,
		}, "    <!--this is a comment-->\n"},
		{3, &html.Node{
			FirstChild: nil,
			Type:       html.TextNode,
			Data:       "this is a text",
			Attr:       nil,
		}, "      this is a text\n"},
	}

	for _, test := range tests {
		out = new(bytes.Buffer) // To capture output

		depth = test.depth
		startElement(test.node)
		got := out.(*bytes.Buffer).String()

		if got != test.want {
			t.Errorf("startElement(%v) prints %q, want %q", test.node, got, test.want)
		}
	}
}

func TestEndElement(t *testing.T) {
	tests := []struct {
		depth int
		node  *html.Node
		want  string
	}{
		{0, &html.Node{
			FirstChild: &html.Node{},
			Type:       html.ElementNode,
			Data:       "html",
			Attr:       nil,
		}, "</html>\n"},
		{1, &html.Node{
			FirstChild: nil,
			Type:       html.ElementNode,
			Data:       "img",
			Attr:       []html.Attribute{{Key: "src", Val: "foo.img"}},
		}, ""},
		{2, &html.Node{
			FirstChild: &html.Node{},
			Type:       html.ElementNode,
			Data:       "a",
			Attr:       []html.Attribute{{Key: "href", Val: "foo.com"}, {Key: "target", Val: "_blank"}},
		}, "    </a>\n"},
		{2, &html.Node{
			FirstChild: nil,
			Type:       html.CommentNode,
			Data:       "this is a comment",
			Attr:       nil,
		}, ""},
		{3, &html.Node{
			FirstChild: nil,
			Type:       html.TextNode,
			Data:       "this is a text",
			Attr:       nil,
		}, ""},
	}

	for _, test := range tests {
		out = new(bytes.Buffer) // To capture output

		depth = test.depth + 1
		endElement(test.node)
		got := out.(*bytes.Buffer).String()

		if got != test.want {
			t.Errorf("endElement(%v) prints %q, want %q", test.node, got, test.want)
		}
	}
}
