package e5_15

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			want: 9,
		},
	}
	for _, test := range tests {
		if got := Max(test.nums...); got != test.want {
			t.Errorf("got = %v, want = %v", got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			want: 0,
		},
	}
	for _, test := range tests {
		if got := Min(test.nums...); got != test.want {
			t.Errorf("got = %v, want = %v", got, test.want)
		}
	}
}
