package chapter2

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	if got := insertSort(arr); !compareIntSlice(got, want) {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
}

func compareIntSlice(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
