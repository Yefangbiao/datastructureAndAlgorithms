package chapter8

// 基数排序
func RadixSort(arr []int) []int {
	//分为正数和负数,对于负数,变成正数最后reverse
	var negatives, noNegatives []int
	for _, num := range arr {
		if num < 0 {
			negatives = append(negatives, -num)
		} else {
			noNegatives = append(noNegatives, num)
		}
	}
	negatives = unsignedRadixSort(negatives)
	p, q := 0, len(negatives)-1
	for p <= q {
		negatives[p], negatives[q] = -negatives[q], -negatives[p]
		p++
		q--
	}
	return append(negatives, unsignedRadixSort(noNegatives)...)
}

func unsignedRadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	maxElement := max(arr)
	for exp := 1; maxElement/exp > 0; exp *= 10 {
		arr = countSort(arr, exp)
	}
	return arr
}

func max(arr []int) int {
	mMax := arr[0]
	for i := 0; i < len(arr); i++ {
		if mMax < arr[i] {
			mMax = arr[i]
		}
	}
	return mMax
}

// 具体排序使用计数排序
func countSort(arr []int, exp int) []int {
	digits := [10]int{}
	for i := 0; i < len(arr); i++ {
		digits[(arr[i]/exp)%10]++
	}
	//统计数组变形
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}
	ans := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		ans[digits[(arr[i]/exp)%10]-1] = arr[i]
		digits[(arr[i]/exp)%10]--
	}
	return ans
}
