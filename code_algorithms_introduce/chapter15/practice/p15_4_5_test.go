package practice

import "testing"

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 4, 3, 3, 6, 7},
			want: []int{1, 2, 2, 2, 3, 4},
		},
	}
	for _, test := range tests {
		c, b := LongestSubsequence(test.nums)

		printLongestSubsequence(test.nums, c, b)
	}
}
