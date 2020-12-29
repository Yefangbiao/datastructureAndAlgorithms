package p1_3

// 写了两个相同的程序,一个使用普通方法,一个使用了Joins
// 使用了11.4节测试相关
// 对echo和echoJoins分别调用1000次
// 结论,明显使用Joins更快

import (
	"testing"
)

func Benchmark1000(b *testing.B) {
	benchmark(b, 1000)
}
func BenchmarkJoins1000(b *testing.B) {
	benchmarkJoins(b, 1000)
}

func benchmark(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		echo()
	}
}

func benchmarkJoins(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		echoJoins()
	}
}
