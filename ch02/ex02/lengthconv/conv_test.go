package lengthconv

import "testing"

func TestFToM(t *testing.T) {
	f := Foot(100)
	want := Meter(30.48)

	got := FToM(f)

	if got != want {
		t.Errorf("FToM(%v) got %v, want %v", f, got, want)
	}
}

func TestMToF(t *testing.T) {
	m := Meter(100)
	want := Foot(328.084)

	got := MToF(m)

	if got != want {
		t.Errorf("MToF(%v) got %v, want %v", m, got, want)
	}
}
