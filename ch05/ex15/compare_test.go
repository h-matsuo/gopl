package compare

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		vals      []int
		want      int
		wantPanic bool
	}{
		{[]int{}, 0, true},
		{[]int{1}, 1, false},
		{[]int{0, 1, -1, 0}, 1, false},
		{[]int{0, -10, math.MinInt64, 0, 20, math.MaxInt64, 200}, math.MaxInt64, false},
	}

	for _, test := range tests {
		func() {
			defer func() {
				p := recover()
				if test.wantPanic && p == nil {
					t.Errorf("max(%v) expects to be panicked, but not panicked", test.vals)
				} else if !test.wantPanic && p != nil {
					t.Errorf("max(%v) expects not to be panicked, but panicked", test.vals)
				}
			}()
			got := max(test.vals...)
			if got != test.want {
				t.Errorf("max(%v) got %d, want %d", test.vals, got, test.want)
			}
		}()
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		vals      []int
		want      int
		wantPanic bool
	}{
		{[]int{}, 0, true},
		{[]int{1}, 1, false},
		{[]int{0, 1, -1, 0}, -1, false},
		{[]int{0, -10, math.MinInt64, 0, 20, math.MaxInt64, 200}, math.MinInt64, false},
	}

	for _, test := range tests {
		func() {
			defer func() {
				p := recover()
				if test.wantPanic && p == nil {
					t.Errorf("min(%v) expects to be panicked, but not panicked", test.vals)
				} else if !test.wantPanic && p != nil {
					t.Errorf("min(%v) expects not to be panicked, but panicked", test.vals)
				}
			}()
			got := min(test.vals...)
			if got != test.want {
				t.Errorf("min(%v) got %d, want %d", test.vals, got, test.want)
			}
		}()
	}
}

func TestAltMax(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{0, 1, -1, 0}, 1},
		{[]int{0, -10, math.MinInt64, 0, 20, math.MaxInt64, 200}, math.MaxInt64},
	}

	for _, test := range tests {
		got := altMax(test.vals[0], test.vals[1:]...)
		if got != test.want {
			t.Errorf("max(%v) got %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestAltMin(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{0, 1, -1, 0}, -1},
		{[]int{0, -10, math.MinInt64, 0, 20, math.MaxInt64, 200}, math.MinInt64},
	}

	for _, test := range tests {
		got := altMin(test.vals[0], test.vals[1:]...)
		if got != test.want {
			t.Errorf("min(%v) got %d, want %d", test.vals, got, test.want)
		}
	}
}
