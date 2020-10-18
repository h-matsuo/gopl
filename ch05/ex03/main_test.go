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
		{"", []string{}},
		{"<html><head><style>NG</style></head><body>a<script>NG</script>b<div>c</div></body></html>", []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.html))

		got := visit(nil, doc)

		if reflect.DeepEqual(got, test.want) {
			t.Errorf("visit() results in %q, want %q", got, test.want)
		}
	}
}
