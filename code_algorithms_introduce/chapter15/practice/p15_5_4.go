package practice

import (
	"math"
)

// 改一行就行
func OptimalBST(p []float64, q []float64, n int) ([][]float64, [][]int) {
	e := make([][]float64, n+2)
	w := make([][]float64, n+2)
	for i := range e {
		e[i] = make([]float64, n+1)
		w[i] = make([]float64, n+1)
	}
	for i := 1; i <= n+1; i++ {
		e[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}
	root := make([][]int, n+1)
	for i := range root {
		root[i] = make([]int, n+1)
		// 改之前初始化一下
		root[i][i] = i
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i][j] = math.MaxInt32
			w[i][j] = w[i][j-1] + p[j] + q[j]
			// 加一行单独判断行=列的情况
			if i == j {
				r := i
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j] = t
					root[i][j] = r
				}
				continue
			}

			// 改下面 这一行
			for r := root[i][j-1]; r <= root[i+1][j]; r++ {
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j] = t
					root[i][j] = r
				}
			}
		}
	}
	return e, root
}
