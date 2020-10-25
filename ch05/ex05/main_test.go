package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	tests := []struct {
		html                  string
		wantWords, wantImages int
	}{
		{"", 0, 0},
		{`
		<html>
			<body>
				<!-- NG: this is a comment -->
				this  is 	a  text   node
				<div>
					<img>
					<p>
						<img>
						<a href="link-a">this is also text node</a>
					</p>
				</div>
				<img>
			</body>
		</html>
		`, 10, 3},
	}

	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.html))

		gotWords, gotImages := countWordsAndImages(doc)

		if gotWords != test.wantWords {
			t.Errorf("countWordsAndImages() reports words: %d, want %d", gotWords, test.wantWords)
		}
		if gotImages != test.wantImages {
			t.Errorf("countWordsAndImages() reports images: %d, want %d", gotImages, test.wantImages)
		}
	}
}
