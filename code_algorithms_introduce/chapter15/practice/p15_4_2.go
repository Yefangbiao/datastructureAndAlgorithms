package practice

import "fmt"

// LCSLength  计算最长公共子序列
func LCSLength(X, Y string) [][]int {
	m := len(X)
	n := len(Y)
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
			} else {
				c[i][j] = c[i][j-1]
			}
		}
	}

	return c
}

func printLCS(c [][]int, X, Y string, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	if X[i-1] == Y[j-1] {
		printLCS(c, X, Y, i-1, j-1)
		fmt.Printf("%c", X[i-1])
	} else {
		// 找三个方向最大的
		if c[i-1][j-1] >= c[i][j-1] && c[i-1][j-1] >= c[i-1][j] {
			printLCS(c, X, Y, i-1, j-1)
		} else if c[i-1][j] >= c[i-1][j-1] && c[i-1][j] >= c[i][j-1] {
			printLCS(c, X, Y, i-1, j)
		} else {
			printLCS(c, X, Y, i, j-1)
		}
	}
}
