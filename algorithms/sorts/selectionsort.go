package sorts

func SelectionSort(arr []int) []int {
	// 选择排序，非递增排序
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}
