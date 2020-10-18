package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		html string
		want map[string]int
	}{
		{"<html><head></head><body></body></html>", map[string]int{"html": 1, "head": 1, "body": 1}},
		{`
		<html>
			<head><link><link></head>
			<body>
				<!-- this is a comment -->
				<a>text</a>
				<div>
					<a>text</a>
					<a>text</a>
					<p><a>text</a></p>
				</div>
			</body>
		</html>
		`, map[string]int{"html": 1, "head": 1, "link": 2, "body": 1, "a": 4, "div": 1, "p": 1}},
	}

	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.html))

		counts := make(map[string]int)
		visit(counts, doc)

		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("visit() results in %q, want %q", counts, test.want)
		}
	}
}
