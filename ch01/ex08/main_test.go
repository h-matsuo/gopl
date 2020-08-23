package main

import "testing"

func TestAddSchemeIfNotSpecified(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"http://example.com", "http://example.com"},
		{"example.com", "http://example.com"},
	}

	for _, test := range tests {
		got := addSchemeIfNotSpecified(test.url)
		if got != test.want {
			t.Errorf("addSchemeIfNotSpecified(%q) got %q, want %q", test.url, got, test.want)
		}
	}
}
