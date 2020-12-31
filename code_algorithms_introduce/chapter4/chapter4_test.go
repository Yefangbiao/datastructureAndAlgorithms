package chapter4

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
		if got := FindMaximumSubArrayBrutal(test.arr); got != test.want {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
		if got := FindMaximumSubArray(test.arr); got != test.want {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
	}
}

func TestSquareMatrixMultiply(t *testing.T) {
	tests := []struct {
		A    [][]int
		B    [][]int
		want [][]int
	}{
		{
			A:    [][]int{{1, 2}, {3, 4}},
			B:    [][]int{{3, 4}, {1, 2}},
			want: [][]int{{5, 8}, {13, 20}},
		},
	}
	for _, test := range tests {
		if got := SquareMatrixMultiply(test.A, test.B); !compareSquare(got, test.want) {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
		if got := StrassenMatrixMultiply(test.A, test.B); !compareSquare(got, test.want) {
			t.Errorf("error, got = %v, want = %v", got, test.want)
		}
	}
}

func BenchmarkFindMaximumSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints := make([]int, i)
		FindMaximumSubArrayBrutal(ints)
	}
}

func BenchmarkFindMaximumSubArrayBrutal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints := make([]int, i)
		FindMaximumSubArrayBrutal(ints)
	}
}

func compareSquare(A, B [][]int) bool {
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A); i++ {
		if len(A[i]) != len(B[i]) {
			return false
		}
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}
	return true
}
