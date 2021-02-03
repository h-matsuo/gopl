package equal

import (
	"testing"
)

const f = 1.0

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{f, f, true},
		{f + threshold/10.0, f, true},
		{f, f + threshold/10.0, true},
		{f + threshold*2, f, false},
		{f, f + threshold*2, false},
		{complex(f, f*2), complex(f, f*2), true},
		{complex(f, f*2+threshold/10.0), complex(f+threshold/10.0, f*2), true},
		{complex(f+threshold/10.0, f*2), complex(f, f*2+threshold/10.0), true},
		{complex(f+threshold*2, f*2+threshold*2), complex(f, f*2), false},
		{complex(f, f*2), complex(f+threshold*2, f*2+threshold*2), false},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
