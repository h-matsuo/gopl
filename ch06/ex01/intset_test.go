package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 65}, 2},
		{[]int{1, 65, 66, 520}, 4},
	}

	for _, test := range tests {
		s := &IntSet{}
		for _, value := range test.values {
			s.Add(value)
		}

		got := s.Len()

		if got != test.want {
			t.Errorf("Len() results in %d, want %d", got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		inits   []int
		removes []int
		want    []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1, 65}, []int{}},
		{[]int{1}, []int{1}, []int{}},
		{[]int{1}, []int{65}, []int{1}},
		{[]int{1, 65}, []int{65}, []int{1}},
		{[]int{1, 65, 66, 520, 620}, []int{520}, []int{1, 65, 66, 620}},
	}

	for _, test := range tests {
		s := &IntSet{}
		for _, value := range test.inits {
			s.Add(value)
		}

		for _, value := range test.removes {
			s.Remove(value)
		}

		for _, value := range test.want {
			if !s.Has(value) {
				t.Errorf("Remove() removes wrong value %d", value)
				break
			}
		}
		if s.Len() != len(test.want) {
			t.Errorf("Remove() results in %s, want %v", s.String(), test.want)
		}

		s.Add(50)
		if !s.Has(50) {
			t.Errorf("Cannot add value after Remove()")
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 65}},
		{[]int{1, 65, 66, 520}},
	}

	for _, test := range tests {
		s := &IntSet{}
		for _, value := range test.values {
			s.Add(value)
		}

		s.Clear()

		if s.Len() != 0 {
			t.Errorf("Clear() fails to clear all values")
		}

		s.Add(50)
		if !s.Has(50) {
			t.Errorf("Cannot add value after Clear()")
		}
	}
}

func TestCopy(t *testing.T) {
	s1 := &IntSet{}

	s1.Copy() // Check empty set can be copied

	s1.Add(1)
	s1.Add(560)
	s2 := s1.Copy()
	if !s2.Has(1) || !s2.Has(560) {
		t.Errorf("Copy() failed")
	}

	s2.Remove(1)
	if !s1.Has(1) || !s1.Has(560) {
		t.Errorf("Copy() affects the original set")
	}
}
