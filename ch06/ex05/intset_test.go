package intset

import (
	"reflect"
	"sort"
	"testing"
)

// Note: This file contains all tests from Exercise 6.1 -- 6.4

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
