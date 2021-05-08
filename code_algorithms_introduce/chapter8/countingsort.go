package chapter8

// 计数排序实现
func sortArray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	ans := make([]int, n)
	lo, hi := nums[0], nums[0]
	for _, num := range nums {
		lo = Min(lo, num)
		hi = Max(hi, num)
	}
	tmp := make([]int, hi-lo+1)

	for _, num := range nums {
		tmp[num-lo]++
	}

	for i := 1; i < len(tmp); i++ {
		tmp[i] += tmp[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		ans[tmp[nums[i]-lo]-1] = nums[i]
		tmp[nums[i]-lo]--
	}

	return ans
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
