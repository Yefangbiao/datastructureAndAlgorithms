package practice

// 练习2.2 冒泡排序

func BubbleSort(arr []int) []int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
					swap = true
				}
			}
		}
	}
	return arr
}
