package chapter15

import "testing"

func TestLcsLength(t *testing.T) {
	tests := []struct {
		X    string
		Y    string
		want int
	}{
		{X: "ABCBDAB", Y: "BDCABA", want: 4},
	}
	for _, test := range tests {
		c, b := LCSLength(test.X, test.Y)
		m := len(test.X)
		n := len(test.Y)
		if c[m][n] != test.want {
			t.Errorf("LCSLength: got = %v, want = %v", c[m][n], test.want)
		}
		printLCS(b, test.X, m, n)
	}
}
