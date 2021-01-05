package e5_15

import "log"

func Max(nums ...int) int {
	if len(nums) == 0 {
		log.Fatal("Max :0 params")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		log.Fatal("Min :0 params")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}
