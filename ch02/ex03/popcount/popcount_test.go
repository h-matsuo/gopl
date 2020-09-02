package popcount

import "testing"

func TestLoopPopCount(t *testing.T) {
	x := uint64(0xff00ff00)
	got := LoopPopCount(x)
	want := 16
	if got != want {
		t.Errorf("LoopPopCount(%d) got %d, want %d", x, got, want)
	}
}

func BenchmarkOriginalPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalPopCount(uint64(i))
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(uint64(i))
	}
}
