package practice

func FindMaximumSubArray(arr []int) int {
	// 第一步首先问题变换
	A := transformQuestion(arr)
	// 类似于动态规划
	dp := make([]int, len(A))
	dp[0] = A[0]
	ans := dp[0]
	for i := 1; i < len(A); i++ {
		dp[i] = max(A[i], dp[i-1]+A[i])
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
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
