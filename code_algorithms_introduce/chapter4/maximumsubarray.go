package chapter4

import "math"

// 4.1 最大子数组问题

// 1.暴力方法
func FindMaximumSubArrayBrutal(arr []int) int {
	ans := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if ans < arr[j]-arr[i] {
				ans = arr[j] - arr[i]
			}
		}
	}
	return ans
}

// 2.课本方法
func FindMaximumSubArray(arr []int) int {
	// 第一步首先问题变换
	A := transformQuestion(arr)
	// 分治求解
	ans := findmaximumsubarray(A)

	return ans
}

func findmaximumsubarray(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	mid := len(A) / 2
	// 求左边的最大值
	leftMax := findmaximumsubarray(A[:mid])
	// 求右边的最大值
	rightMax := findmaximumsubarray(A[mid:])
	// 求中间的最大值
	crossMax := findmaxcrossingsubarray(A)
	return max(leftMax, rightMax, crossMax)

}

func findmaxcrossingsubarray(A []int) int {
	leftSum := math.MinInt32
	sum := 0
	for i := len(A) / 2; i >= 0; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	rightSum := math.MinInt32
	sum = 0
	for i := len(A)/2 + 1; i < len(A); i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

func transformQuestion(arr []int) []int {
	A := make([]int, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		A[i-1] = arr[i] - arr[i-1]
	}
	return A
}

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if m < nums[i] {
			m = nums[i]
		}
	}
	return m
}
