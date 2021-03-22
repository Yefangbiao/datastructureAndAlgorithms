package chapter6

import (
	"fmt"
	"testing"
)

func TestSortByHeap(t *testing.T) {

	var tests = []struct {
		nums    []int
		compare bool
		want    []int
	}{
		{
			nums:    []int{1, 2, 3},
			compare: false,
			want:    []int{3, 2, 1},
		},
	}
	for _, test := range tests {
		got := SortByHeap(test.nums, len(test.nums), test.compare)
		fmt.Println(got)
	}
}

func TestPQ(t *testing.T) {
	hp := newHeap([]int{}, 0, false)
	hp.Push(3)
	hp.Push(1)
	fmt.Println(hp.Pop())
}
