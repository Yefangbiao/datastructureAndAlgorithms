package practice

// ActivitySelector 选择最晚开始的活动  迭代算法
func ActivitySelector(s []int, f []int) (A []int) {
	n := len(s)
	A = append(A, n-1)
	k := n - 1
	for i := n - 2; i >= 1; i-- {
		if f[i] < s[k] {
			A = append(A, i)
			k = i
		}
	}
	return
}
