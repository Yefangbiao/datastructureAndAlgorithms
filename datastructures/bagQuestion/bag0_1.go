package bagQuestion

import "fmt"

// 0-1背包问题
// 假设并不需要把背包装满

// Goods商品
type Goods struct {
	Weight int
	Value  int
}

// MaxValue 求解0-1背包,输入物品数: n，背包容量: c，输入 n 个物品的重量、价值
// 动态规划方法
func MaxValue(n, c int, goods []Goods) int {

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}
	// 初始化
	if goods[0].Weight <= c {
		dp[0][goods[0].Weight] = goods[0].Value
	}
	fmt.Println(dp[0])
	// 动态规划
	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j <= c; j++ {
			dp[i][j] = dp[i-1][j]
			if goods[i].Weight <= c && j-goods[i].Weight >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-goods[i].Weight]+goods[i].Value)
			}
		}
		fmt.Println(dp[i])
	}
	for j := 0; j <= c; j++ {
		if ans < dp[n-1][j] {
			ans = dp[n-1][j]
		}
	}
	return ans
}

// MaxValue2 回溯or枚举解决0-1背包问题
func MaxValue2(n, c int, goods []Goods) int {
	res := 0
	dfs(n, c, 0, goods, &res, 0, 0)
	return res
}

// dfs 输入物品数: n，背包容量: c，cNow 现在的背包容量,输入 n 个物品的重量、价值
func dfs(n, c, cNow int, goods []Goods, res *int, start int, val int) {
	if cNow > c {
		return
	}
	if val > *res {
		*res = val
	}
	for i := start; i < n; i++ {
		dfs(n, c, cNow+goods[i].Weight, goods, res, i+1, val+goods[i].Value)
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
