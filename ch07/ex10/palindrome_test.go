package palindrome

import (
	"fmt"
	"sort"
	"testing"
)

func main() {
	s := []string{"a", "b", "c", "c", "b", "a"}
	fmt.Println(IsPalindrome(sort.StringSlice(s)))
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abb", false},
		{"abba", true},
		{"わたしまけましたわ", true}, // 私、負けましたわ
	}

	for _, test := range tests {
		s := []string{}
		for _, r := range test.s {
			s = append(s, string(r))
		}

		got := IsPalindrome(sort.StringSlice(s))

		if got != test.want {
			t.Errorf("IsPalindrome(%q) reports palindrome?: %t, want %t", test.s, got, test.want)
		}
	}
}
