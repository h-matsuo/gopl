package sexpr

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type X struct {
		b1, b2 bool
		f      float32
		z      complex64
		c      chan int
		fun    func()
		i      interface{}
	}
	x := X{
		b1:  true,
		b2:  false,
		f:   3 / 10,
		z:   1 + 2i,
		c:   make(chan int),
		fun: func() { return },
		i:   []int{1, 2, 3},
	}

	// Encode it
	data, err := Marshal(x)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
}
