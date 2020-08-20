package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/h-matsuo/gopl/ch01/ex04/counter"
)

func TestMainFromStdin(t *testing.T) {
	input := "foo\nfoo"
	args := []string{"cmd-name"}
	want := fmt.Sprintf("2\t[%s]\tfoo\n", os.Stdin.Name())

	in = io.Reader(bytes.NewBufferString(input))
	out = new(bytes.Buffer) // To capture output
	osArgs = args

	main()

	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("{ input: %q, args: %q } got output %q, want %q", input, args, got, want)
	}
}

func TestMainFromFiles(t *testing.T) {
	file1, file2 := "foo\n", "foo\n"
	fileName1 := filepath.Join(os.TempDir(), "file1")
	fileName2 := filepath.Join(os.TempDir(), "file2")
	args := []string{"cmd-name", fileName1, fileName2}
	want1 := fmt.Sprintf("2\t[%s %s]\tfoo\n", fileName1, fileName2)
	want2 := fmt.Sprintf("2\t[%s %s]\tfoo\n", fileName2, fileName1)

	out = new(bytes.Buffer) // To capture output
	osArgs = args

	ioutil.WriteFile(fileName1, []byte(file1), os.ModePerm)
	defer os.Remove(fileName1)
	ioutil.WriteFile(fileName2, []byte(file2), os.ModePerm)
	defer os.Remove(fileName2)

	main()

	got := out.(*bytes.Buffer).String()
	if got != want1 && got != want2 {
		t.Errorf("{ args: %q, file1: %q, file2: %q } got output %q, want ( %q or %q )", args, file1, file2, got, want1, want2)
	}
}

func TestCountLinesFromStdin(t *testing.T) {
	input := "foo"
	wantLine := input
	wantFileName := os.Stdin.Name()

	r := bytes.NewBufferString(input)
	lc := counter.NewLineCounter()

	countLines(r, lc)

	gotLine := lc.ToSlice()[0].Line()
	if gotLine != wantLine {
		t.Errorf("Input %q from stdin got line %q, want %q", input, gotLine, wantLine)
	}

	gotFileName := lc.ToSlice()[0].FileNames()[0]
	if gotFileName != wantFileName {
		t.Errorf("Input %q from stdin got fileName %q, want %q", input, gotFileName, wantFileName)
	}
}

func TestCountLinesFromFile(t *testing.T) {
	input := "foo"
	fileName := filepath.Join(os.TempDir(), "file")
	wantLine := input
	wantFileName := fileName

	ioutil.WriteFile(fileName, []byte(input), os.ModePerm)
	defer os.Remove(fileName)

	r, _ := os.Open(fileName)
	lc := counter.NewLineCounter()

	countLines(r, lc)

	gotLine := lc.ToSlice()[0].Line()
	if gotLine != wantLine {
		t.Errorf("Input %q from file got line %q, want %q", input, gotLine, wantLine)
	}

	gotFileName := lc.ToSlice()[0].FileNames()[0]
	if gotFileName != wantFileName {
		t.Errorf("Input %q from file got fileName %q, want %q", input, gotFileName, wantFileName)
	}
}

func TestShowResultWithoutOutput(t *testing.T) {
	input := "foo" // Only 1 line -- nothing expected to be printed!
	args := []string{"cmd-name"}

	in = io.Reader(bytes.NewBufferString(input))
	out = new(bytes.Buffer) // To capture output
	osArgs = args

	main()

	got := out.(*bytes.Buffer).String()
	if len(got) > 0 {
		t.Errorf("Input single line %q got output %q, want nothing", input, got)
	}
}
