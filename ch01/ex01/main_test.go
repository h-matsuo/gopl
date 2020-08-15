package main

import (
	"bytes"
	"testing"
)

func TestPrintStringsWithoutArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	osArgs := []string{"cmd-name"} // Only command name w/o command line args

	printStrings(buf, osArgs)

	if buf.String() != "cmd-name" {
		t.Errorf("printStrings(%q) // => %q", osArgs, buf.String())
	}
}

func TestPrintStringsWithSomeArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	osArgs := []string{"cmd-name", "foo", "bar"}

	printStrings(buf, osArgs)

	if buf.String() != "cmd-name foo bar" {
		t.Errorf("printStrings(%q) // => %q", osArgs, buf.String())
	}
}
