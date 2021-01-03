package e4_4

// Rotate return a slice
// rotate the slice in k
// for example [1,2,3,4,5] k=2
// output: [3,4,5,1,2]
func Rotate(arr []int, k int) []int {
	reverse(arr[:k])
	reverse(arr[k:len(arr)])
	reverse(arr)
	return arr
}

func reverse(p []int) {
	for i, j := 0, len(p)-1; i <= j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}
