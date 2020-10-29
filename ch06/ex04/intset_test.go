package intset

import (
	"reflect"
	"sort"
	"testing"
)

func TestElements(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{2, 101}},
		{[]int{1, 2, 101}},
	}

	for _, test := range tests {
		s := &IntSet{}
		sort.Ints(test.values)
		for _, value := range test.values {
			s.Add(value)
		}

		got := s.Elements()
		sort.Ints(got)

		if !reflect.DeepEqual(got, test.values) {
			t.Errorf("TestElements() results in %v, want %v", got, test.values)
		}
	}
}
