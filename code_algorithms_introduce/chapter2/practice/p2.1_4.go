package practice

// 练习2.1-4

func aAddB(A, B []int) []int {
	res := make([]int, len(A)+1)
	carry := 0
	for i := len(A) - 1; i >= 0; i-- {
		res[i+1] = (A[i] + B[i] + carry) % 2
		A[i] += B[i]
		if A[i]+B[i]+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	res[0] = carry

	return res
}
