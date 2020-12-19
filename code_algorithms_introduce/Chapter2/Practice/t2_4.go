package Practice

// 练习2.4

func CountInversions(arr []int) (int, []int) {
	if len(arr) < 2 {
		return 0, arr
	}

	mid := len(arr) / 2
	left, l := CountInversions(arr[:mid])
	right, r := CountInversions(arr[mid:])
	mergeNum, merge := mergeInversions(l, r)
	return mergeNum + left + right, merge
}

func mergeInversions(left, right []int) (int, []int) {
	merge := make([]int, len(left)+len(right))
	ans := 0
	i, j := len(left)-1, len(right)-1
	for i >= 0 && j >= 0 {
		if left[i] > right[j] {
			ans += j + 1
			merge[i+j+1] = left[i]
			i--
		} else {
			merge[i+j+1] = right[j]
			j--
		}
	}
	for i >= 0 {
		merge[i+j+1] = left[i]
		i--
	}
	for j >= 0 {
		merge[i+j+1] = right[j]
		j--
	}
	return ans, merge
}
