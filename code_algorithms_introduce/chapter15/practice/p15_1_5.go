package practice

// CalFib 计算第n个斐波那契数
func CalFib(n int) int {
	if n <= 2 {
		return 1
	}
	fib := make([]int, n+1)
	fib[1] = 1
	fib[2] = 1
	for i := 3; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
