package chapter15

import "fmt"

// LCSLength  计算最长公共子序列
func LCSLength(X, Y string) ([][]int, [][]int) {
	m := len(X)
	n := len(Y)
	c := make([][]int, m+1)
	// 2代表X[i]=Y[i]是LCS的一个元素,1代表向上，0代表向前
	b := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
		b[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = 2
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = 1
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = 0
			}
		}
	}

	return c, b
}

func printLCS(b [][]int, X string, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == 2 {
		printLCS(b, X, i-1, j-1)
		fmt.Printf("%c", X[i-1])
	} else if b[i][j] == 1 {
		printLCS(b, X, i-1, j)
	} else {
		printLCS(b, X, i, j-1)
	}
}
