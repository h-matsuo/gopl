package popcount

func BitClearPopCount(x uint64) int {
	count := 0
	for x != 0 { // Wait for all bits to be cleared
		x = x & (x - 1) // x&(x-1) clears the least significant bit where 1 is set
		count++
	}
	return count
}
