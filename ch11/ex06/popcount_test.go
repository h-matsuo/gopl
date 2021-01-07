package popcount

import (
	"testing"
)

// To avoid compiler's optimization on benchmark code
var (
	output int
)

func benchmarkPopCount(b *testing.B, input uint64) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += PopCount(input)
	}
	output = tmp
}

func benchmarkLSBPopCount(b *testing.B, input uint64) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += LSBPopCount(input)
	}
	output = tmp
}

func benchmarkBitClearPopCount(b *testing.B, input uint64) {
	var tmp int // To avoid compiler's optimization on benchmark code
	for i := 0; i < b.N; i++ {
		tmp += BitClearPopCount(input)
	}
	output = tmp
}

func BenchmarkPopCount_0(b *testing.B)         { benchmarkPopCount(b, 0x0) }
func BenchmarkLSBPopCount_0(b *testing.B)      { benchmarkLSBPopCount(b, 0x0) }
func BenchmarkBitClearPopCount_0(b *testing.B) { benchmarkBitClearPopCount(b, 0x0) }

func BenchmarkPopCount_FFFF(b *testing.B)         { benchmarkPopCount(b, 0xFFFF) }
func BenchmarkLSBPopCount_FFFF(b *testing.B)      { benchmarkLSBPopCount(b, 0xFFFF) }
func BenchmarkBitClearPopCount_FFFF(b *testing.B) { benchmarkBitClearPopCount(b, 0xFFFF) }

func BenchmarkPopCount_FFFFFFFF(b *testing.B)         { benchmarkPopCount(b, 0xFFFFFFFF) }
func BenchmarkLSBPopCount_FFFFFFFF(b *testing.B)      { benchmarkLSBPopCount(b, 0xFFFFFFFF) }
func BenchmarkBitClearPopCount_FFFFFFFF(b *testing.B) { benchmarkBitClearPopCount(b, 0xFFFFFFFF) }

func BenchmarkPopCount_FFFFFFFFFFFFFFFF(b *testing.B)    { benchmarkPopCount(b, 0xFFFFFFFFFFFFFFFF) }
func BenchmarkLSBPopCount_FFFFFFFFFFFFFFFF(b *testing.B) { benchmarkLSBPopCount(b, 0xFFFFFFFFFFFFFFFF) }
func BenchmarkBitClearPopCount_FFFFFFFFFFFFFFFF(b *testing.B) {
	benchmarkBitClearPopCount(b, 0xFFFFFFFFFFFFFFFF)
}
