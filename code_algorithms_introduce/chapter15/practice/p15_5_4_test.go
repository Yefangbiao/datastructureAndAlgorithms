package practice

import (
	"fmt"
	"testing"
)

func TestOptimalBST(t *testing.T) {
	tests := []struct {
		p []float64
		q []float64
		n int
	}{
		{
			p: []float64{0, 0.15, 0.10, 0.05, 0.10, 0.20},
			q: []float64{0.05, 0.10, 0.05, 0.05, 0.05, 0.10},
			n: 5,
		},
		{
			p: []float64{0, 0.04, 0.06, 0.08, 0.02, 0.10, 0.12, 0.14},
			q: []float64{0.06, 0.06, 0.06, 0.06, 0.05, 0.05, 0.05, 0.05},
			n: 7,
		},
	}
	for _, test := range tests {
		e, root := OptimalBST(test.p, test.q, test.n)
		for i := 1; i <= test.n+1; i++ {
			for j := 0; j <= test.n; j++ {
				fmt.Printf("%.3f ", e[i][j])
			}
			fmt.Println()
		}

		for i := 1; i <= test.n; i++ {
			for j := 0; j <= test.n; j++ {
				fmt.Print(root[i][j], " ")
			}
			fmt.Println()
		}

	}
}
