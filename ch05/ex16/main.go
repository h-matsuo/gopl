package main

import (
	"strings"
)

func join(sep string, elems ...string) string {
	var result strings.Builder
	for i, elem := range elems {
		if i > 0 {
			result.WriteString(sep)
		}
		result.WriteString(elem)
	}
	return result.String()
}
