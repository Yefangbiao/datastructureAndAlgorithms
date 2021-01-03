package e4_3

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		arr  [10]int
		want [10]int
	}{
		{
			arr:  [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		if Reverse(&test.arr); test.arr != test.want {
			t.Errorf("error, got = %v, want = %v\n", test.arr, test.want)
		}
	}
}
