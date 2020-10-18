package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	template := `
	<html>
		<head>
			<link rel="stylesheet" integrity="dummy" href="link-stylesheet-1">
			<link href="link-stylesheet-2" rel="stylesheet" media="all">
			<link rel="canonical" href="NG">
		</head>
		<body>
			<!-- this is a comment -->
			<div>
				<img src="link-img" alt="text">
				<p><a href="link-a">text</a></p>
			</div>
			<script src="link-script" async></script>
			<script>/* NG */</script>
		</body>
	</html>
	`
	tests := []struct {
		flags *config
		html  string
		want  []string
	}{
		{&config{}, "", []string{}},
		{&config{a: true}, template, []string{"link-a"}},
		{&config{img: true}, template, []string{"link-img"}},
		{&config{script: true}, template, []string{"link-script"}},
		{&config{stylesheet: true}, template, []string{"link-stylesheet-1", "link-stylesheet-2"}},
		{&config{a: true, img: true, script: true, stylesheet: true}, template, []string{"link-stylesheet-1", "link-stylesheet-2", "link-img", "link-a", "link-script"}},
	}

	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.html))

		got := visit(nil, doc)

		if reflect.DeepEqual(got, test.want) {
			t.Errorf("visit() results in %q, want %q", got, test.want)
		}
	}
}
