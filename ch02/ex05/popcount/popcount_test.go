package popcount

import "testing"

func TestBitClearPopCount(t *testing.T) {
	x := uint64(0xff00ff00)
	got := BitClearPopCount(x)
	want := 16
	if got != want {
		t.Errorf("BitClearPopCount(%d) got %d, want %d", x, got, want)
	}
}

func BenchmarkBitClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitClearPopCount(uint64(i))
	}
}
