package sorts

import "math/rand"

func QuickSort(arr []int) []int {
	return quickSort(arr)
}

func quickSort(arr []int) []int {
	// 三取样切分
	if len(arr) <= 1 {
		return arr
	}
	x := arr[rand.Intn(len(arr))]
	lowPart := make([]int, 0, len(arr))
	equalPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	for _, num := range arr {
		if num < x {
			lowPart = append(lowPart, num)
		} else if num == x {
			equalPart = append(equalPart, num)
		} else {
			highPart = append(highPart, num)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, equalPart...)
	lowPart = append(lowPart, highPart...)

	return arr
}
