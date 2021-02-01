package chapter16

// GreedyActivitySelector 迭代的贪心选择
func GreedyActivitySelector(s []int, f []int) (A []int) {
	n := len(s)
	A = append(A, 1)
	k := 1
	for i := 2; i < n; i++ {
		if s[i] >= f[k] {
			A = append(A, i)
			k = i
		}
	}
	return
}

// RecursiveActivitySelector 递归的贪心选择
func RecursiveActivitySelector(s []int, f []int, k, n int) []int {
	m := k + 1
	for m < n && s[m] < f[k] {
		m++
	}
	res := make([]int, 0, n)
	if m < n {
		res = append(res, m)
		return append(res, RecursiveActivitySelector(s, f, m, n)...)
	} else {
		return []int{}
	}
}
