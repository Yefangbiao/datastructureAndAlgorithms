package practice

func QuickSort712(A []int) {
	quickSort712(A, 0, len(A)-1)
}

// quickSort712 快速排序
func quickSort712(A []int, p, r int) {
	if p < r {
		q := partition712(A, p, r)
		quickSort712(A, p, q-1)
		quickSort712(A, q+1, r)
	}
}

func partition712(A []int, p, r int) int {
	x := A[r]
	i := p - 1

	flag := true

	for j := p; j < r; j++ {
		if A[j] != x {
			flag = false
		}
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	if flag {
		return (p + r) / 2
	}
	return i + 1
}
