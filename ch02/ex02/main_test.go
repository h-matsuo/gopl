package main

import (
	"io"
	"strings"
	"testing"
)

func TestGetInput(t *testing.T) {
	tests := []struct {
		in           io.Reader
		osArgs       []string
		wantQuantity float64
		wantUnit     string
	}{
		{nil, []string{"cmd-name", "10.0", "kg"}, 10.0, "kg"},
		{strings.NewReader("10.0 kg"), []string{"cmd-name"}, 10.0, "kg"},
	}

	for _, test := range tests {
		in = test.in
		osArgs = test.osArgs

		gotQuantity, gotUnit := getInput()

		if gotQuantity != test.wantQuantity {
			t.Errorf("getInput() with Stdin %q, os.Args %q got quantity %f, want %f", test.in, test.osArgs, gotQuantity, test.wantQuantity)
		}
		if gotUnit != test.wantUnit {
			t.Errorf("getInput() with Stdin %q, os.Args %q got unit %q, want %q", test.in, test.osArgs, gotUnit, test.wantUnit)
		}
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		quantity float64
		unit     string
		want     string
	}{
		{100, "°C", "212.000000 °F"},
		{32, "°F", "0.000000 °C"},
		{100, "ft", "30.480000 m"},
		{100, "m", "328.084000 ft"},
		{100, "lb", "45.359237 kg"},
		{100, "kg", "220.460000 lb"},
	}

	for _, test := range tests {
		got := convert(test.quantity, test.unit)

		if got != test.want {
			t.Errorf("convert(%f, %q) got %q, want %q", test.quantity, test.unit, got, test.want)
		}
	}
}
