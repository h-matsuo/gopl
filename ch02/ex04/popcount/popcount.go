package popcount

func LSBPopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		x >>= 1
		count += int(x & uint64(0b1))
	}
	return count
}
