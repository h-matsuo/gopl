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
		want []string
	}{
		{"", nil},
		{`
		<html>
			<head><link href="NG"></head>
			<body>
				<!-- this is a comment -->
				<a href="https://link1.com">text</a>
				<div>
					<a class="my-class" href="/link2" id="my-id">text</a>
					<a href="link3" rel="noreferrer">text</a>
					<p><a href="#link4" target="_blank">text</a></p>
				</div>
			</body>
		</html>
		`, []string{"https://link1.com", "/link2", "link3", "#link4"}},
	}

	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.html))

		got := visit(nil, doc)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("visit() returns %q, want %q", got, test.want)
		}
	}
}
