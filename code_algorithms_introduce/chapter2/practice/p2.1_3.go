package practice

// 练习2.1-3
func linearSearch(nums []int, v int) int {
	for i, num := range nums {
		if v == num {
			return i
		}
	}
	// nil使用-1来代替
	return -1
}
