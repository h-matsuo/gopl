package intset

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

func toString(m map[int]bool) string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}

	buf.WriteByte('}')
	return buf.String()
}

func TestHas(t *testing.T) {
	tests := []struct {
		inputs []int
		target int
	}{
		{[]int{}, 1},
		{[]int{1}, 1},
		{[]int{1}, 2},
		{[]int{1, 144, 9}, 144},
		{[]int{1, 144, 9, 9, 144, 3}, 9},
	}
	for _, test := range tests {
		var got IntSet
		want := make(map[int]bool)
		for _, n := range test.inputs {
			got.Add(n)
			want[n] = true
		}
		if got.Has(test.target) != want[test.target] {
			t.Errorf("Add %q to IntSet and Has(%d) got %t, want %t", test.inputs, test.target, got.Has(test.target), want[test.target])
		}
	}
}

func TestAdd(t *testing.T) {
	tests := [][]int{
		[]int{},
		[]int{1},
		[]int{1, 144, 9},
		[]int{1, 144, 9, 9, 144, 3},
	}
	for _, test := range tests {
		var got IntSet
		want := make(map[int]bool)
		for _, n := range test {
			got.Add(n)
			want[n] = true
		}
		if got.String() != toString(want) {
			t.Errorf("Add %q to IntSet and String() got %q, want %q", test, got.String(), toString(want))
		}
	}
}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		x, y []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{2}},
		{[]int{1, 1024}, []int{144, 1}},
	}
	for _, test := range tests {
		var got IntSet
		want := make(map[int]bool)
		for _, n := range test.x {
			got.Add(n)
			want[n] = true
		}
		var y IntSet
		for _, n := range test.y {
			y.Add(n)
			want[n] = true
		}
		got.UnionWith(&y)
		if got.String() != toString(want) {
			t.Errorf("Union %q with %q and String() got %q, want %q", test.x, test.y, got.String(), toString(want))
		}
	}
}
