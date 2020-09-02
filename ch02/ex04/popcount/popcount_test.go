package popcount

import "testing"

func TestLSBPopCount(t *testing.T) {
	x := uint64(0xff00ff00)
	got := LSBPopCount(x)
	want := 16
	if got != want {
		t.Errorf("LSBPopCount(%d) got %d, want %d", x, got, want)
	}
}

func BenchmarkLSBPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LSBPopCount(uint64(i))
	}
}
