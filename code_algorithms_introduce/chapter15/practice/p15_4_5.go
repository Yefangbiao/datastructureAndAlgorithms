package practice

import "fmt"

func LongestSubsequence(nums []int) ([]int, []int) {
	c := make([]int, len(nums))
	b := make([]int, len(nums))
	for i := range c {
		c[i] = 1
		b[i] = i
	}
	b[0] = -1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				c[i] = c[j] + 1
				b[i] = j
				break
			}
		}
	}
	return c, b
}

func printLongestSubsequence(nums []int, c, b []int) {
	max := len(c) - 1
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] > c[max] {
			max = i
		}
	}
	stack := make([]int, 0, len(nums)-1)
	for max >= 0 {
		stack = append(stack, nums[max])
		max = b[max]
	}
	for len(stack) > 0 {
		fmt.Print(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
}
