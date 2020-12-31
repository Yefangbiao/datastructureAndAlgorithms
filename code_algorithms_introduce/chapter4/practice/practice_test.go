package practice

import "testing"

func TestFindMaximumSubArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97},
			want: 43,
		},
		{
			arr:  []int{10, 11, 7, 10, 6},
			want: 3,
		},
		{
			arr:  []int{-5, -2, -3},
			want: 3,
		},
	}
	for _, test := range tests {
		if got := FindMaximumSubArray(test.arr); got != test.want {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
	}
}
