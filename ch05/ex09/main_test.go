package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"", func(_ string) string { panic("NG") }, ""},
		{"$this\tis\n$sample", func(s string) string { return strings.ToUpper(s) }, "THIS\tis\nSAMPLE"},
		{
			"  $マルチバイト文字 も  $単語間のスペース数   も正しく扱える必要があります ",
			func(_ string) string { return "foo" },
			"  foo も  foo   も正しく扱える必要があります ",
		},
	}

	for _, test := range tests {
		got := expand(test.s, test.f)

		if got != test.want {
			t.Errorf("expand() results in %v, want %v", got, test.want)
		}
	}
}
