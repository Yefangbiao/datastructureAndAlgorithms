package chapter8

import "sort"

// 桶排序
// 用于排序正数
func BucketSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	minVal := minInSlice(nums...)
	maxVal := maxInSlice(nums...)

	// 桶的长度
	d := Max(1, maxVal-minVal/(n-1))

	// 桶的个数
	bucketCount := (maxVal - minVal/d) + 1
	buckets := make([][]int, bucketCount)

	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	// 放入每个桶
	for _, num := range nums {
		buckets[(num-minVal)/d] = append(buckets[(num-minVal)/d], num)
	}
	for i := 0; i < len(buckets); i++ {
		// 桶内使用其他排序
		sort.Ints(buckets[i])
		if i > 0 {
			buckets[0] = append(buckets[0], buckets[i]...)
		}

	}
	return buckets[0]
}

func minInSlice(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func maxInSlice(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
