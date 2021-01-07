package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCount(t *testing.T) {
	type result struct {
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}
	tests := []struct {
		input string
		want  result
	}{
		{"", result{map[rune]int{}, [utf8.UTFMax + 1]int{}, 0}},
		{"a\na", result{map[rune]int{'a': 2, '\n': 1}, [utf8.UTFMax + 1]int{0, 3, 0, 0, 0}, 0}},
		{"a©あ🍺", result{map[rune]int{'a': 1, '©': 1, 'あ': 1, '🍺': 1}, [utf8.UTFMax + 1]int{0, 1, 1, 1, 1}, 0}},
	}

	for _, test := range tests {
		in := strings.NewReader(test.input)
		gotCounts, gotUtflen, gotInvalid := count(in)
		if !reflect.DeepEqual(gotCounts, test.want.counts) {
			t.Errorf("count(%q) got counts: %v, want: %v", test.input, gotCounts, test.want.counts)
		}
		if !reflect.DeepEqual(gotUtflen, test.want.utflen) {
			t.Errorf("count(%q) got utflen: %v, want: %v", test.input, gotUtflen, test.want.utflen)
		}
		if gotInvalid != test.want.invalid {
			t.Errorf("count(%q) got invalid: %q, want: %q", test.input, gotInvalid, test.want.invalid)
		}
	}
}
