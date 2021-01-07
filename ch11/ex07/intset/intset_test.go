package intset

import (
	"math/rand"
	"testing"
)

// To avoid compiler's optimization on benchmark code
var (
	outIntSet IntSet
	outMap    map[int]bool
)

// Benchmark for Add()

func benchmarkIntSetAdd(b *testing.B, num, maxNumber int) {
	// Generate data set
	b.StopTimer()
	rng := rand.New(rand.NewSource(0)) // Reset seed for same value every time
	data := make([]int, 0, num)
	for i := 0; i < num; i++ {
		data = append(data, rng.Intn(maxNumber))
	}
	var x IntSet
	// Start benchmark
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			x.Add(d)
		}
	}
	outIntSet = x
}

func benchmarkMapAdd(b *testing.B, num, maxNumber int) {
	// Generate data set
	b.StopTimer()
	rng := rand.New(rand.NewSource(0)) // Reset seed for same value every time
	data := make([]int, 0, num)
	for i := 0; i < num; i++ {
		data = append(data, rng.Intn(maxNumber))
	}
	x := make(map[int]bool)
	// Start benchmark
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			x[d] = true
		}
	}
	outMap = x
}

func BenchmarkIntSetAdd_num1_max1(b *testing.B) { benchmarkIntSetAdd(b, 1, 1) }
func BenchmarkMapAdd_num1_max1(b *testing.B)    { benchmarkMapAdd(b, 1, 1) }

func BenchmarkIntSetAdd_num1000_max1000(b *testing.B) { benchmarkIntSetAdd(b, 1000, 1000) }
func BenchmarkMapAdd_num1000_max1000(b *testing.B)    { benchmarkMapAdd(b, 1000, 1000) }

func BenchmarkIntSetAdd_num1000_max10000(b *testing.B) { benchmarkIntSetAdd(b, 10000, 10000) }
func BenchmarkMapAdd_num10000_max10000(b *testing.B)   { benchmarkMapAdd(b, 10000, 10000) }

// Benchmark for UnionWith()

func benchmarkIntSetUnionWith(b *testing.B, num, maxNumber int) {
	// Generate data set
	b.StopTimer()
	rng := rand.New(rand.NewSource(0)) // Reset seed for same value every time
	var x, y IntSet
	for i := 0; i < num; i++ {
		x.Add(rng.Intn(maxNumber))
		y.Add(rng.Intn(maxNumber))
	}
	// Start benchmark
	b.StartTimer()
	x.UnionWith(&y)
}

func benchmarkMapUnionWith(b *testing.B, num, maxNumber int) {
	// Generate data set
	b.StopTimer()
	rng := rand.New(rand.NewSource(0)) // Reset seed for same value every time
	x, y := make(map[int]bool), make(map[int]bool)
	for i := 0; i < num; i++ {
		x[rng.Intn(maxNumber)] = true
		y[rng.Intn(maxNumber)] = true
	}
	// Start benchmark
	b.StartTimer()
	for d := range y {
		x[d] = true
	}
}

func BenchmarkIntSetUnionWith_num1_max1(b *testing.B) { benchmarkIntSetUnionWith(b, 1, 1) }
func BenchmarkMapUnionWith_num1_max1(b *testing.B)    { benchmarkMapUnionWith(b, 1, 1) }

func BenchmarkIntSetUnionWith_num1000_max1000(b *testing.B) {
	benchmarkIntSetUnionWith(b, 1000, 1000)
}
func BenchmarkMapUnionWith_num1000_max1000(b *testing.B) { benchmarkMapUnionWith(b, 1000, 1000) }

func BenchmarkIntSetUnionWith_num1000_max10000(b *testing.B) {
	benchmarkIntSetUnionWith(b, 10000, 10000)
}
func BenchmarkMapUnionWith_num10000_max10000(b *testing.B) { benchmarkMapUnionWith(b, 10000, 10000) }
