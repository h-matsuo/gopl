package popcount

import "sync"

var pc [256]byte

var initPcOnce = make([]sync.Once, 256)

func lookupPc(idx byte) byte {
	initPcOnce[idx].Do(func() {
		if idx == 0 {
			pc[idx] = pc[idx/2] + byte(idx&1)
		} else {
			pc[idx] = lookupPc(idx/2) + byte(idx&1)
		}
	})
	return pc[idx]
}

func PopCount(x uint64) int {
	return int(lookupPc(byte(x>>(0*8))) +
		lookupPc(byte(x>>(1*8))) +
		lookupPc(byte(x>>(2*8))) +
		lookupPc(byte(x>>(3*8))) +
		lookupPc(byte(x>>(4*8))) +
		lookupPc(byte(x>>(5*8))) +
		lookupPc(byte(x>>(6*8))) +
		lookupPc(byte(x>>(7*8))))
}
