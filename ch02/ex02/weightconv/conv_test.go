package weightconv

import "testing"

func TestPToK(t *testing.T) {
	p := Pound(100)
	want := Kilogram(45.359237)

	got := PToK(p)

	if got != want {
		t.Errorf("PToK(%v) got %v, want %v", p, got, want)
	}
}

func TestKToP(t *testing.T) {
	k := Kilogram(100)
	want := Pound(220.46)

	got := KToP(k)

	if got != want {
		t.Errorf("KToP(%v) got %v, want %v", k, got, want)
	}
}
