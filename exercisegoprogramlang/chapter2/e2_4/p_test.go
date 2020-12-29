package p2_4

import "testing"

func Benchmark1000(b *testing.B) {
	benchmarkPopCount(b, 1000)
}
func BenchmarkPopCountRight1000(b *testing.B) {
	benchmarkPopCountRight(b, 1000)
}

func benchmarkPopCount(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func benchmarkPopCountRight(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCountRight(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func TestPopCount(t *testing.T) {
	PopCountRight(0x123)
}
