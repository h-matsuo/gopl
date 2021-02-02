package display

import (
	"testing"
)

// This test ensures that the program terminates without crashing.
func TestMapKeyOfStructOrArray(t *testing.T) {
	type MArray map[[2]int]int
	mArray := MArray{
		[2]int{1, 2}: -1,
		[2]int{3, 4}: -2,
	}
	Display("mArray", mArray)

	type Struct struct {
		a int
		b string
	}
	type MStruct map[Struct]int
	mStruct := MStruct{
		Struct{a: 1, b: "foo"}: -1,
		Struct{a: 2, b: "bar"}: -2,
	}

	Display("mStruct", mStruct)
}
