package Practice

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	want := []int{6, 5, 4, 3, 2, 1}
	if got := insertSort(arr); !compareIntSlice(got, want) {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
}

func TestLinearSearch(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	want := 3
	if got := linearSearch(arr, 6); got != want {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
	want = -1
	if got := linearSearch(arr, 10); got != want {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
}

func TestAAddB(t *testing.T) {
	A := []int{1, 1, 1, 1}
	B := []int{1, 1, 1, 1}
	want := []int{1, 1, 1, 1, 0}
	if got := aAddB(A, B); !compareIntSlice(got, want) {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}

}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	if got := selectionSort(arr); !compareIntSlice(got, want) {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	want := 2
	if got := BinarySearch(arr, 3); got != want {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
	want = -1
	if got := BinarySearch(arr, 7); got != want {
		fmt.Printf("error! got :%v, want :%v", got, want)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	if got := MergeSort(arr); !compareIntSlice(got, want) {
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
