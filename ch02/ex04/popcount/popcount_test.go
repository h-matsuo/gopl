package popcount

import (
	"math/bits"
	"testing"
)

func TestLSBPopCount(t *testing.T) {
	x := uint64(0xf123456789abcdef)
	got := LSBPopCount(x)
	want := bits.OnesCount64(x)
	if got != want {
		t.Errorf("LSBPopCount(%d) got %d, want %d", x, got, want)
	}
}

// To avoid compiler's optimization on benchmark code
var (
	input  = uint64(0xf123456789abcdef)
	output int
)

func BenchmarkLSBPopCount(b *testing.B) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += LSBPopCount(input)
	}
	output = tmp
}
