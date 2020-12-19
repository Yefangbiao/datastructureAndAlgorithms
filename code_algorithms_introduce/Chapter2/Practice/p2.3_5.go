package Practice

// 练习2.3-5 二分查找
// 返回要查找的数字位置,如果没有找到返回-1
func BinarySearch(arr []int, v int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {

		mid := (lo + hi) / 2
		if arr[mid] == v {
			return mid
		}
		if arr[mid] < v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return -1
}
