package chapter15

import (
	"fmt"
	"math"
)

// 矩阵链乘法
// 此过程假定矩阵Ai的规模为Pi-1 X Pi(i=1,2,...,n)
// 输入时一个序列p=<p0,p1,p2,...,pn>,其长度为p.length=n+1
// 过程调用一个辅助表m保存代价m[i,j],另一个辅助表s记录最优值m[i,j]对应分割点k

// MatrixChainOrder 计算矩阵链乘的最小代价
func MatrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p) - 1
	m := make([][]int, n+1)
	s := make([][]int, n+1)
	// 初始化
	for i := 0; i <= n; i++ {
		m[i] = make([]int, n+1)
		s[i] = make([]int, n+1)
	}
	//l is the chain length
	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt32
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}
	return m, s
}

func PrintOptimalParens(s [][]int, i, j int) {
	if i == j {
		fmt.Printf("A%d", i)
	} else {
		fmt.Print("(")
		PrintOptimalParens(s, i, s[i][j])
		PrintOptimalParens(s, s[i][j]+1, j)
		fmt.Print(")")
	}
}
