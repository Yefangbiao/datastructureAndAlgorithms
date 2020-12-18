package Practice

// 使用二分法查找插入排序优于归并排序的时候的值
// !! 由于存在精度损失,这应该是一个近似值

import (
	"fmt"
	"math"
)

func findN_1_2_2() int {
	lo, hi := 1, 100
	for lo < hi {
		mid := (lo + hi) / 2
		if calc1_1_2_2(mid) {
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	return lo
}

func calc1_1_2_2(x int) bool {
	res := 8*x*x - 64*x*(int(math.Log2(float64(x))))
	fmt.Println(x, res)
	if res > 0 {
		return true
	}
	return false
}
