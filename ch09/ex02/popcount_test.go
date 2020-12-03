package popcount

import (
	"math/bits"
	"testing"
)

func TestPopCount(t *testing.T) {
	x := uint64(0xf123456789abcdef)
	got := PopCount(x)
	want := bits.OnesCount64(x)
	if got != want {
		t.Errorf("PopCount(%d) got %d, want %d", x, got, want)
	}
}
