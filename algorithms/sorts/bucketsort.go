package sorts

import "math"

func BucketSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	max, min := math.MinInt32, math.MaxInt32
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	//桶的间距
	d := maxBucket(1, (max-min)/(len(arr)-1))
	//注意`[1,1,1]`d为0的情况
	//桶的个数
	bucketSize := (max-min)/d + 1
	//桶
	buckets := make([][]int, bucketSize)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	//入桶
	for _, num := range arr {
		location := (num - min) / d
		buckets[location] = append(buckets[location], num)
	}

	//桶内用计数排序
	for i := range buckets {
		buckets[i] = countSortBucket(buckets[i])
	}
	//返回
	ans := make([]int, 0, len(arr))
	for i := range buckets {
		ans = append(ans, buckets[i]...)
	}
	return ans
}

func maxBucket(nums ...int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func countSortBucket(arr []int) []int {
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
