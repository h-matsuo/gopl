package tempconv

import "testing"

func TestCToF(t *testing.T) {
	c := BoilingC
	want := Fahrenheit(212)

	got := CToF(c)

	if got != want {
		t.Errorf("CToF(%v) got %v, want %v", c, got, want)
	}
}

func TestCToK(t *testing.T) {
	c := BoilingC
	want := Kelvin(373.15)

	got := CToK(c)

	if got != want {
		t.Errorf("CToK(%v) got %v, want %v", c, got, want)
	}
}

func TestFToC(t *testing.T) {
	f := Fahrenheit(32)
	want := FreezingC

	got := FToC(f)

	if got != want {
		t.Errorf("FToC(%v) got %v, want %v", f, got, want)
	}
}

func TestFToK(t *testing.T) {
	f := Fahrenheit(32)
	want := Kelvin(273.15)

	got := FToK(f)

	if got != want {
		t.Errorf("FToK(%v) got %v, want %v", f, got, want)
	}
}

func TestKToC(t *testing.T) {
	k := AbsoluteZeroK
	want := Celsius(-273.15)

	got := KToC(k)

	if got != want {
		t.Errorf("KToC(%v) got %v, want %v", k, got, want)
	}
}

func TestKToF(t *testing.T) {
	k := Kelvin(273.15)
	want := Fahrenheit(32)

	got := KToF(k)

	if got != want {
		t.Errorf("KToF(%v) got %v, want %v", k, got, want)
	}
}
