package practice

import (
	"fmt"
	"testing"
)

func TestActivitySelector(t *testing.T) {
	tests := []struct {
		s    []int
		f    []int
		want int
	}{
		{
			s:    []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12},
			f:    []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16},
			want: 4,
		},
	}
	for _, test := range tests {
		got1 := ActivitySelector(test.s, test.f)
		if len(got1) != test.want {
			t.Errorf("GreedyActivitySelector: got = %v, want = %v", len(got1), test.want)
		}
		fmt.Println(got1)
	}
}
