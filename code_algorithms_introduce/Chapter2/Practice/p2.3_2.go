package Practice

// 练习2.3-2
// 归并排序
func MergeSort(arr []int) []int {
	return merge(arr)
}

func merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := merge(arr[:mid])
	right := merge(arr[mid:])
	i, j := 0, 0
	var ans []int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}
	for i < len(left) {
		ans = append(ans, left[i])
		i++
	}
	for j < len(right) {
		ans = append(ans, right[j])
		j++
	}
	return ans
}
