package counter

import (
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	tests := []struct {
		l    *LineCounter
		want []*Counter
	}{
		{NewLineCounter(), make([]*Counter, 0)},
		{
			func() *LineCounter {
				l := NewLineCounter()
				l.AddLine("line", "fileName")
				return l
			}(),
			func() []*Counter {
				l := NewLineCounter()
				l.AddLine("line", "fileName")
				return []*Counter{l.counterByLine["line"]}
			}(),
		},
	}

	for _, test := range tests {
		got := test.l.ToSlice()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v.ToSlice() got %+v, want %+v", test.l, got, test.want)
		}
	}
}

func TestAddLine(t *testing.T) {
	tests := []struct {
		l        *LineCounter
		line     string
		fileName string
		want     *LineCounter
	}{
		{
			NewLineCounter(),
			"line",
			"fileName",
			func() *LineCounter {
				l := NewLineCounter()
				c := &Counter{
					line:        "line",
					occurrences: 0,
					fileNames:   make(map[string]struct{}),
				}
				c.increment("fileName")
				l.counterByLine = map[string]*Counter{"line": c}
				return l
			}(),
		},
		{
			func() *LineCounter {
				l := NewLineCounter()
				c := &Counter{
					line:        "line",
					occurrences: 0,
					fileNames:   make(map[string]struct{}),
				}
				c.increment("fileName")
				l.counterByLine = map[string]*Counter{"line": c}
				return l
			}(),
			"line",
			"fileName",
			func() *LineCounter {
				l := NewLineCounter()
				c := &Counter{
					line:        "line",
					occurrences: 1,
					fileNames:   make(map[string]struct{}),
				}
				c.increment("fileName")
				l.counterByLine = map[string]*Counter{"line": c}
				return l
			}(),
		},
	}

	for _, test := range tests {
		test.l.AddLine(test.line, test.fileName)
		got := test.l
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v.AddLine(%q, %q) got %+v, want %+v", test.l, test.line, test.fileName, got, test.want)
		}
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		c        *Counter
		fileName string
		want     *Counter
	}{
		{
			&Counter{
				line:        "line",
				occurrences: 0,
				fileNames:   make(map[string]struct{}),
			},
			"fileName",
			&Counter{
				line:        "line",
				occurrences: 1,
				fileNames:   map[string]struct{}{"fileName": struct{}{}},
			},
		},
		{
			&Counter{
				line:        "line",
				occurrences: 1,
				fileNames:   map[string]struct{}{"fileName": struct{}{}},
			},
			"fileName",
			&Counter{
				line:        "line",
				occurrences: 2,
				fileNames:   map[string]struct{}{"fileName": struct{}{}},
			},
		},
	}

	for _, test := range tests {
		test.c.increment(test.fileName)
		got := test.c
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v.increment(%q) got %+v, want %+v", test.c, test.fileName, got, test.want)
		}
	}
}

func TestFileNames(t *testing.T) {
	tests := []struct {
		c    *Counter
		want []string
	}{
		{
			&Counter{
				line:        "line",
				occurrences: 0,
				fileNames:   make(map[string]struct{}),
			},
			[]string{},
		},
		{
			&Counter{
				line:        "line",
				occurrences: 0,
				fileNames:   map[string]struct{}{"fileName": struct{}{}},
			},
			[]string{"fileName"},
		},
	}

	for _, test := range tests {
		got := test.c.FileNames()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%+v.FileNames() got %+v, want %+v", test.c, got, test.want)
		}
	}
}
