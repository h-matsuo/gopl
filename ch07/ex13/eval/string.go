package eval

import (
	"fmt"
	"strconv"
)

///// Solutions to Exercise 7.13 ////////////////////

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("( %c%s )", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("( %s %c %s )", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("pow( %s, %s )", c.args[0].String(), c.args[1].String())
	case "sin":
		return fmt.Sprintf("sin( %s )", c.args[0].String())
	case "sqrt":
		return fmt.Sprintf("sqrt( %s )", c.args[0].String())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
