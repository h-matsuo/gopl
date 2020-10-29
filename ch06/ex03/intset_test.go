package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{1, 2, 101, 102}, []int{2, 101}, []int{2, 101}},
		{[]int{2, 101}, []int{1, 2, 101, 102}, []int{2, 101}},
		{[]int{1, 2, 101, 102}, []int{1, 102}, []int{1, 102}},
		{[]int{1, 102}, []int{1, 2, 101, 102}, []int{1, 102}},
	}

	for _, test := range tests {
		a, b := &IntSet{}, &IntSet{}
		for _, value := range test.a {
			a.Add(value)
		}
		for _, value := range test.b {
			b.Add(value)
		}

		a.IntersectWith(b)

		for _, value := range test.want {
			if !a.Has(value) {
				t.Errorf("IntSet doesn't contain %d after IntersectWith()", value)
				break
			}
		}
		if a.Len() != len(test.want) {
			t.Errorf("IntersectWith() results in %s, want %v", a.String(), test.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{1}, []int{}},
		{[]int{1, 2, 101, 102}, []int{2, 101}, []int{1, 102}},
		{[]int{2, 101}, []int{1, 2, 101, 102}, []int{}},
		{[]int{1, 2, 101, 102}, []int{1, 102}, []int{2, 101}},
		{[]int{1, 102}, []int{1, 2, 101, 102}, []int{}},
	}

	for _, test := range tests {
		a, b := &IntSet{}, &IntSet{}
		for _, value := range test.a {
			a.Add(value)
		}
		for _, value := range test.b {
			b.Add(value)
		}

		a.DifferenceWith(b)

		for _, value := range test.want {
			if !a.Has(value) {
				t.Errorf("IntSet doesn't contain %d after DifferenceWith()", value)
				break
			}
		}
		if a.Len() != len(test.want) {
			t.Errorf("DifferenceWith() results in %s, want %v", a.String(), test.want)
		}
	}
}

func TestSymmetricWith(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{1}, []int{}},
		{[]int{1, 2, 101, 102}, []int{2, 101}, []int{1, 102}},
		{[]int{2, 101}, []int{1, 2, 101, 102}, []int{1, 102}},
		{[]int{1, 2, 101, 102}, []int{1, 102}, []int{2, 101}},
		{[]int{1, 102}, []int{1, 2, 101, 102}, []int{2, 101}},
	}

	for _, test := range tests {
		a, b := &IntSet{}, &IntSet{}
		for _, value := range test.a {
			a.Add(value)
		}
		for _, value := range test.b {
			b.Add(value)
		}

		a.SymmetricWith(b)

		for _, value := range test.want {
			if !a.Has(value) {
				t.Errorf("IntSet doesn't contain %d after SymmetricWith()", value)
				break
			}
		}
		if a.Len() != len(test.want) {
			t.Errorf("SymmetricWith() results in %s, want %v", a.String(), test.want)
		}
	}
}
