package main

import (
	"net/http"
	"reflect"
	"testing"
)

func fillDefaultValues(c *lissajousConfig) *lissajousConfig {
	defaultConfig := newLissajousConfig()
	if c.cycles == 0 {
		c.cycles = defaultConfig.cycles
	}
	if c.res == 0.0 {
		c.res = defaultConfig.res
	}
	if c.size == 0 {
		c.size = defaultConfig.size
	}
	if c.nframes == 0 {
		c.nframes = defaultConfig.nframes
	}
	if c.delay == 0 {
		c.delay = defaultConfig.delay
	}
	return c
}

func TestParseQuery(t *testing.T) {
	tests := []struct {
		url          string
		want         *lissajousConfig
		willGotError bool
	}{
		{"/", newLissajousConfig(), false},
		{"/?cycles=100", fillDefaultValues(&lissajousConfig{cycles: 100}), false},
		{"/?res=0.1", fillDefaultValues(&lissajousConfig{res: 0.1}), false},
		{"/?size=100", fillDefaultValues(&lissajousConfig{size: 100}), false},
		{"/?nframes=100", fillDefaultValues(&lissajousConfig{nframes: 100}), false},
		{"/?delay=100", fillDefaultValues(&lissajousConfig{delay: 100}), false},
		{"/?cycles=100&res=0.1", fillDefaultValues(&lissajousConfig{cycles: 100, res: 0.1}), false},
		{"/?cycles=a", nil, true},
		{"/?res=0.a", nil, true},
	}

	for _, test := range tests {
		r, _ := http.NewRequest("GET", test.url, nil)

		got, errGot := parseQuery(r)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseQuery() with url %q got %+v, want %+v", test.url, got, test.want)
		}
		if (errGot != nil) != test.willGotError {
			t.Errorf("parseQuery() with url %q got error %q, expects error?: %t", test.url, errGot, test.willGotError)
		}
	}
}
