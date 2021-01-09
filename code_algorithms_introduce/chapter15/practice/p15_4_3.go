package practice

// 递归的LCS
// 严格的来说不是O(MN),在特殊情况下，但是就先这样吧
func RecursiveLCSLength(X, Y string) [][]int {
	m := len(X)
	n := len(Y)
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}
	recursiveLCSLength(X, Y, m, n, c)
	return c
}

func recursiveLCSLength(X, Y string, i, j int, c [][]int) {
	if i == 0 || j == 0 {
		return
	}
	if c[i][j] > 0 {
		return
	}
	if X[i-1] == Y[j-1] {
		recursiveLCSLength(X, Y, i-1, j-1, c)
		c[i][j] = c[i-1][j-1] + 1
	} else {
		recursiveLCSLength(X, Y, i-1, j, c)
		recursiveLCSLength(X, Y, i, j-1, c)
		if c[i-1][j] >= c[i][j-1] {
			c[i][j] = c[i-1][j]
		} else {
			c[i][j] = c[i][j-1]
		}
	}
}
