package sexpr

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type X struct {
		B1, B2 bool
		F      float32
		Z      complex64
		C      chan int
		A      map[struct{ X int }]<-chan []string
		// I interface{}
	}
	m := make(map[struct{ X int }]<-chan []string)
	m[struct{ X int }{1}] = make(chan []string, 5)
	m[struct{ X int }{2}] = make(<-chan []string, 10)
	x := X{
		B1: true,
		B2: false,
		F:  3 / 10,
		Z:  1 + 2i,
		C:  make(chan int),
		A:  m,
		// I: []int{1, 2, 3},
	}

	// Encode it
	data, err := Marshal(x)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	if err := Unmarshal(data, &x); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", x)
}
