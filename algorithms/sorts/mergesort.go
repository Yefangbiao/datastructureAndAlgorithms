package sorts

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	ans := make([]int, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans[i+j] = left[i]
			i++
		} else {
			ans[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		ans[i+j] = left[i]
		i++
	}
	for j < len(right) {
		ans[i+j] = right[j]
		j++
	}
	return ans
}
