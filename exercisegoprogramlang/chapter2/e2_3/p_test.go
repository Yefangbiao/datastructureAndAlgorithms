package p2_3

import (
	"testing"
)

func Benchmark1000(b *testing.B) {
	benchmarkPopCount(b, 1000)
}
func BenchmarkPopCountFor1000(b *testing.B) {
	benchmarkPopCountFor(b, 1000)
}

func benchmarkPopCount(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func benchmarkPopCountFor(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCountFor(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func TestPopCount(t *testing.T) {
	PopCountFor(0x123)
}
