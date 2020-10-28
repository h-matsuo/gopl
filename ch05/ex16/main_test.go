package main

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		sep   string
		elems []string
	}{
		{"", []string{}},
		{" ", []string{"a"}},
		{" ", []string{"a", "b"}},
		{", ", []string{"ab", "bc", "ca"}},
	}

	for _, test := range tests {
		got := join(test.sep, test.elems...)
		want := strings.Join(test.elems, test.sep)
		if got != want {
			t.Errorf("join(%q, %q) got %q, want %q", test.sep, test.elems, got, want)
		}
	}
}
