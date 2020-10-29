// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	///// Solutions to Exercise 7.15 ////////////////////
	if _, ok := env[v]; !ok {
		fmt.Printf("Found new variable: %q\n", v)
		fmt.Printf("Enter corresponding value: ")
		reader := bufio.NewReader(os.Stdin)
		valueStr, _ := reader.ReadString('\n')
		valueStr = strings.Replace(valueStr, "\n", "", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(err)
		}
		env[v] = value
	}
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

///// Solutions to Exercise 7.14 ////////////////////

func (r repetition) Eval(env Env) float64 {
	numRepetition := int(r.y.Eval(env))
	if numRepetition < 1 {
		return 0
	}
	x := r.x.Eval(env)
	result := x
	for i := 0; i < numRepetition-1; i++ {
		switch r.op {
		case '+':
			result += x
		case '-':
			result -= x
		case '*':
			result *= x
		case '/':
			result /= x
		default:
			panic(fmt.Sprintf("unsupported binary operator: %q", r.op))
		}
	}
	return result
}

//!-Eval2
