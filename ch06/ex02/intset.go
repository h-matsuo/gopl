package intset

import (
	"bytes"
	"fmt"

	"github.com/h-matsuo/gopl/ch02/ex05/popcount"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

///// Solutions to Exercise 6.1 ////////////////////

// Len returns the length of the set.
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += popcount.BitClearPopCount(word)
	}
	return len
}

// Remove removes the value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
	// Remove trailing elements if empty
	for i := len(s.words); i > 0; i-- {
		if s.words[i-1] == 0 {
			s.words = s.words[:i-1]
		} else {
			break
		}
	}
}

// Clear removes all values from the set.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns the copy of the set.
func (s *IntSet) Copy() *IntSet {
	t := &IntSet{}
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return t
}

///// Solutions to Exercise 6.2 ////////////////////

// AddAll adds all values from xList.
func (s *IntSet) AddAll(xList ...int) {
	for _, value := range xList {
		s.Add(value)
	}
}
