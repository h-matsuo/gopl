package popcount

import (
	"math/bits"
	"testing"
)

func TestBitClearPopCount(t *testing.T) {
	x := uint64(0xf123456789abcdef)
	got := BitClearPopCount(x)
	want := bits.OnesCount64(x)
	if got != want {
		t.Errorf("BitClearPopCount(%d) got %d, want %d", x, got, want)
	}
}

// To avoid compiler's optimization on benchmark code
var (
	input  = uint64(0xf123456789abcdef)
	output int
)

func BenchmarkBitClearPopCount(b *testing.B) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += BitClearPopCount(input)
	}
	output = tmp
}
