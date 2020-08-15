package main

import (
	"bytes"
	"testing"
)

func TestPrintStringsWithoutArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	osArgs := []string{"cmd-name"} // Only command name w/o command line args

	printStrings(buf, osArgs)

	if buf.String() != "" {
		t.Errorf("printStrings(%q) // => %q", osArgs, buf.String())
	}
}

func TestPrintStringsWithSomeArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	osArgs := []string{"cmd-name", "foo", "bar"}

	printStrings(buf, osArgs)

	if buf.String() != "(0, foo)\n(1, bar)\n" {
		t.Errorf("printStrings(%q) // => %q", osArgs, buf.String())
	}
}
