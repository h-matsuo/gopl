package circular

import (
	"fmt"
	"testing"
)

func TestIsCircular(t *testing.T) {
	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2, cyclePtr3 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr3

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	type CycleStruct struct {
		p *CycleStruct
	}
	cycleStruct1 := CycleStruct{}
	cycleStruct1.p = &cycleStruct1
	cycleStruct2 := CycleStruct{}

	type CycleMap map[int]*CycleMap
	cycleMap := CycleMap{}
	cycleMap[0] = &cycleMap

	for _, test := range []struct {
		v    interface{}
		want bool
	}{
		{nil, false},
		{1, false},
		{cyclePtr1, true},
		{cyclePtr2, false},
		{cyclePtr3, false},
		{cycleSlice, true},
		{cycleStruct1, true},
		{cycleStruct2, false},
		{cycleMap, true},
	} {
		if IsCircular(test.v) != test.want {
			t.Errorf("IsCircular(%v) = %t", test.v, !test.want)
		}
	}
}

func Example_cycle() {
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(IsCircular(a)) // "true"
	fmt.Println(IsCircular(b)) // "true"
	fmt.Println(IsCircular(c)) // "true"
	d, e := &link{value: "d"}, &link{value: "e"}
	d.tail = e
	fmt.Println(IsCircular(d)) // "false"

	// Output:
	// true
	// true
	// true
	// false
}
