package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.4
func LSBPopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 0b1)
		x >>= 1
	}
	return count
}

// Exercise 2.5
func BitClearPopCount(x uint64) int {
	count := 0
	for x != 0 { // Wait for all bits to be cleared
		x = x & (x - 1) // x&(x-1) clears the least significant bit where 1 is set
		count++
	}
	return count
}
