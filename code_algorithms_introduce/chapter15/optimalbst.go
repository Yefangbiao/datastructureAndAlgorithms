package chapter15

import (
	"fmt"
	"math"
)

// OptimalBST 计算最优搜索二叉树
// p的有效输入是p1,p2...pn. q的有效输入是q0,q1...qn. n是规模
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
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i][j] = math.MaxInt32
			w[i][j] = w[i][j-1] + p[j] + q[j]
			for r := i; r <= j; r++ {
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

// constructOptimalBST 输出最优二叉树
// 迭代有点难，我用递归
func constructOptimalBST(root [][]int, i, j int, str string) {

	if j >= i {
		fmt.Printf("%s-k%d\n", str, root[i][j])
		constructOptimalBST(root, i, root[i][j]-1, str+"-left")
		constructOptimalBST(root, root[i][j]+1, j, str+"-right")
	} else {
		fmt.Printf("%s-d%d\n", str, j)
	}
}
