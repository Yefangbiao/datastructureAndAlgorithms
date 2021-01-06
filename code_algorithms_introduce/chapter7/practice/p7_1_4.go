package practice

func QuickSortReverse(A []int) {
	quickSortReverse(A, 0, len(A)-1)
}

// quickSortReverse 快速排序
func quickSortReverse(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSortReverse(A, p, q-1)
		quickSortReverse(A, q+1, r)
	}
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] >= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
