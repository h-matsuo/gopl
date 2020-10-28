package main

import (
	"math"
)

func max(vals ...int) int {
	if len(vals) < 1 {
		panic("no values specified")
	}
	max := math.MinInt64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) < 1 {
		panic("no values specified")
	}
	min := math.MaxInt64
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func altMax(val1 int, vals ...int) int {
	max := val1
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func altMin(val1 int, vals ...int) int {
	min := val1
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
