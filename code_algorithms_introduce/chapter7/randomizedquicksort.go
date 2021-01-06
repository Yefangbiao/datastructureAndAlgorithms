package chapter7

import (
	"math/rand"
	"time"
)

func RandomizedQuickSort(A []int) {
	randomizedQuickSort(A, 0, len(A)-1)
}

func randomizedQuickSort(A []int, p, r int) {
	if p < r {
		q := randomizedPartition(A, p, r)
		randomizedQuickSort(A, p, q-1)
		randomizedQuickSort(A, p+1, r)
	}
}

func randomizedPartition(A []int, p, r int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(r-p+1) + p
	A[i], A[r] = A[r], A[i]
	return partition(A, p, r)
}
