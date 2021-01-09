package chapter15

import (
	"fmt"
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	m, s := MatrixChainOrder([]int{5, 10, 3, 12, 5, 50, 6})
	for i := range m {
		fmt.Println(m[i])
	}
	PrintOptimalParens(s, 1, len(s)-1)
}
