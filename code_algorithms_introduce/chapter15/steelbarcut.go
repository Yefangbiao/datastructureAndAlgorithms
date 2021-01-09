package chapter15

import (
	"fmt"
	"math"
)

// 钢条切割问题

// CutRod 直接的自顶向下递归实现
// p: 钢条长度的价格 n:长度
func CutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := math.MinInt32
	for i := 1; i <= n; i++ {
		q = MAX(q, p[i]+CutRod(p, n-i))
	}
	return q
}

// MemoizedCutRod  带备忘录的自顶向下法
func MemoizedCutRod(p []int, n int) int {
	r := make([]int, n+1)
	for i := 1; i < len(r); i++ {
		r[i] = math.MinInt32
	}
	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}
	q := math.MinInt32
	for i := 1; i <= n; i++ {
		q = MAX(q, p[i]+memoizedCutRodAux(p, n-i, r))
	}
	return q
}

// BottomUpCutRod 自底向上
func BottomUpCutRod(p []int, n int) int {
	r := make([]int, n+1)
	for i := 1; i < len(r); i++ {
		r[i] = math.MinInt32
	}
	for i := 1; i < len(r); i++ {
		for j := 1; j <= i; j++ {
			r[i] = MAX(r[i], p[j]+r[i-j])
		}
	}
	return r[n]
}

// ExtendedBottomUpCutRod 返回最优收益和对应切割方案
func ExtendedBottomUpCutRod(p []int, n int) ([]int, []int) {
	r := make([]int, n+1)
	s := make([]int, n+1)
	for i := 1; i < len(r); i++ {
		q := math.MinInt32
		for j := 1; j <= i; j++ {
			if q < p[j]+r[i-j] {
				q = p[j] + r[i-j]
				s[i] = j
			}
			r[i] = q
		}
	}
	return r, s
}

// PrintCutRodSolution 打印最优收益和切割方案
func PrintCutRodSolution(p []int, n int) {
	r, s := ExtendedBottomUpCutRod(p, n)
	fmt.Println("r: ", r)
	fmt.Println("s: ", s)
}

func MAX(x, y int) int {
	if x > y {
		return x
	}
	return y
}
