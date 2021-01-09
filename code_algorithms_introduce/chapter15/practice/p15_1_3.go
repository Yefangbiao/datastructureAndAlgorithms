package practice

import "math"

// 解决方法，在计算两段价值的时候加上-c
// BottomUpCutRod 自底向上
func BottomUpCutRod(p []int, n int, c int) int {
	r := make([]int, n+1)
	for i := 1; i < len(r); i++ {
		r[i] = math.MinInt32
	}
	for i := 1; i < len(r); i++ {
		for j := 1; j <= i; j++ {
			r[i] = MAX(r[i], p[j]+r[i-j]-c)
		}
	}
	return r[n]
}

func MAX(x, y int) int {
	if x > y {
		return x
	}
	return y
}
