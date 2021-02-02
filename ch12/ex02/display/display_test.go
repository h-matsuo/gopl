package display

import (
	"testing"
)

const limitDepth = 10

// This test ensures that the program terminates without crashing.
func TestCycle(t *testing.T) {
	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	Display("p", p, limitDepth)

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	Display("m", m, limitDepth)

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	Display("s", s, limitDepth)

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c, limitDepth)
}
