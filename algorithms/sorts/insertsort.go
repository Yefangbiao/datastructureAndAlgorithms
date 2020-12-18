// insertion sort

package sorts

func InsertionSort(arr []int) []int {
	// 将arr按照非降序排序
	for i := 1; i < len(arr); i++ {
		currentNum := arr[i]
		j := i
		for ; j > 0 && arr[j-1] >= currentNum; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = currentNum
	}
	return arr
}
