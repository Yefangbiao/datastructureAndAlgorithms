package practice

import "testing"

func TestRecursiveLCSLength(t *testing.T) {
	tests := []struct {
		X    string
		Y    string
		want int
	}{
		{X: "ABCBDAB", Y: "BDCABA", want: 4},
	}
	for _, test := range tests {
		c := RecursiveLCSLength(test.X, test.Y)
		m := len(test.X)
		n := len(test.Y)
		if c[m][n] != test.want {
			t.Errorf("LCSLength: got = %v, want = %v", c[m][n], test.want)
		}
	}
}
