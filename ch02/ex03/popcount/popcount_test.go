package popcount

import (
	"math/bits"
	"testing"
)

func TestLoopPopCount(t *testing.T) {
	x := uint64(0xf123456789abcdef)
	got := LoopPopCount(x)
	want := bits.OnesCount64(x)
	if got != want {
		t.Errorf("LoopPopCount(%d) got %d, want %d", x, got, want)
	}
}

// To avoid compiler's optimization on benchmark code
var (
	input  = uint64(0xf123456789abcdef)
	output int
)

func BenchmarkOriginalPopCount(b *testing.B) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += OriginalPopCount(input)
	}
	output = tmp
}

func BenchmarkLoopPopCount(b *testing.B) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += LoopPopCount(input)
	}
	output = tmp
}
