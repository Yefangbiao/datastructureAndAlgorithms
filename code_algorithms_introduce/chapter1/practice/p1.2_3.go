package practice

// 使用二分法查找100n^2优于2^n的n
// !!不存在精度损失,这是一个准确值

import (
	"math"
)

func findN_1_2_3() int {
	lo, hi := 1, 100
	for lo < hi {
		mid := (lo + hi) / 2
		if calc1_2_3(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func calc1_2_3(x int) bool {
	res := 100*x*x - int(math.Pow(2, float64(x)))
	if res > 0 {
		return true
	}
	return false
}
