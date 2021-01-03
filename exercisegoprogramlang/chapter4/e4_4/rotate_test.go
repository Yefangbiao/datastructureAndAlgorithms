package e4_4

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
		k    int
	}{
		{
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			k:    2,
		},
	}
	for _, test := range tests {
		if got := Rotate(test.arr, test.k); !compareSlice(got, test.want) {
			t.Errorf("error, got = %v, want = %v\n", got, test.want)
		}
	}
}

func compareSlice(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if a2[i] != v {
			return false
		}
	}
	return true
}
