package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		osArgs []string
		want   string
	}{
		{[]string{"cmd-name"}, ""},
		{[]string{"cmd-name", "foo", "bar"}, "(0, foo)\n(1, bar)\n"},
	}

	for _, test := range tests {
		out = new(bytes.Buffer) // To capture output
		osArgs = test.osArgs

		main()

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("Args %q echoes %q, want %q", test.osArgs, got, test.want)
		}
	}
}
