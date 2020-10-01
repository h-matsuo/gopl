package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%q\n", string(distinctSpace([]byte("  Hello, \u0085 \t\t\n\fworld! \u00a0"))))
}

func distinctSpace(b []byte) []byte {
	isPrevSpace := false
	for i := 0; i < len(b); i++ {
		isCurrentSpace, needsPeek := isSpace(b, i)
		if isCurrentSpace {
			b[i] = ' '
			if needsPeek {
				copy(b[i+1:], b[i+2:])
				b = b[:len(b)-1]
			}
		}
		if isPrevSpace && isCurrentSpace {
			copy(b[i:], b[i+1:])
			b = b[:len(b)-1]
			i--
		}
		isPrevSpace = isCurrentSpace
	}
	return b
}

// isSpace reports whether the character pointed by `idx` is a space character.
// Second return value becomes `true` if the space takes 2 bytes such as
// `'\u0085'` (NEL), `'\u00A0'` (NBSP).
func isSpace(b []byte, idx int) (bool, bool) {
	switch b[idx] {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true, false
	}
	if idx < len(b)-1 && b[idx] == 0xc2 {
		switch b[idx+1] {
		case 0x85, 0xa0:
			return true, true
		}
	}
	return false, false
}
