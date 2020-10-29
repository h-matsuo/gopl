package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{nil},
		{[]int{}},
		{[]int{1}},
		{[]int{500}},
		{[]int{1, 65, 66, 620}},
	}

	for _, test := range tests {
		s := &IntSet{}

		s.AddAll(test.values...)

		for _, value := range test.values {
			if !s.Has(value) {
				t.Errorf("IntSet doesn't contain %d after AddAll()", value)
				break
			}
		}
	}
}
