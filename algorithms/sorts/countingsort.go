package sorts

import (
	"math"
)

func CountingSort(arr []int) []int {
	if len(arr) < 1 {
		// 注意空数组特判
		return arr
	}
	max, min := math.MinInt32, math.MaxInt32
	//1.统计最大值和最小值
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	//2.建立一个区间数组
	count := make([]int, max-min+1, max-min+1)
	for _, num := range arr {
		count[num-min]++
	}
	//3.统计变形
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	//4.倒序遍历
	res := make([]int, len(arr), len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
	}
	return res
}
