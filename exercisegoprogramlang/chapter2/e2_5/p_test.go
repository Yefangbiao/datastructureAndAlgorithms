package p2_5

import "testing"

func Benchmark1000(b *testing.B) {
	benchmarkPopCount(b, 1000)
}
func BenchmarkPopCountAnd1000(b *testing.B) {
	benchmarkPopCountAnd(b, 1000)
}

func benchmarkPopCount(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func benchmarkPopCountAnd(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCountAnd(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func TestPopCount(t *testing.T) {
	PopCountAnd(0x123)
}
