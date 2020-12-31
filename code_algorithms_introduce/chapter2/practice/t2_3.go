package practice

// 练习2.3

// Horner Horner方法
func Horner(arr []int, x int) int {
	y := 0
	for i := len(arr) - 1; i >= 0; i-- {
		y = arr[i] + x*y
	}
	return y
}

func CalMulti(arr []int, x int) int {
	y := 0
	for k := 0; k < len(arr); k++ {
		temp := 1
		for i := 0; i < k; i++ {
			temp = temp * x
		}
		y = y + arr[k]*temp
	}
	return y
}
