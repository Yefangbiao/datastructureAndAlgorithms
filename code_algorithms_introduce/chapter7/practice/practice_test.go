package practice

import (
	"testing"
)

type sortTest struct {
	input    []int
	expected []int
	name     string
}

func TestQuickSortReverse(t *testing.T) {
	sortTests := []sortTest{
		//Sorted slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "Sorted Unsigned"},
		//Reversed slice
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "Reversed Unsigned"},
	}
	for _, test := range sortTests {
		QuickSortReverse(test.input)
		if pos, b := compareSlices(test.input, test.expected); b != true {
			t.Errorf("error, pos = %d", pos)
		}
	}
}

func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}
