package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte{'x'})
	// c1是[32]int 类型
	ans := 0
	for i := 0; i < len(c1); i++ {
		ans += PopCountAnd(c1[i])
	}
	fmt.Printf("digits of c1 is %d\n", ans)

}

// 根据练习2.5统计位数
func PopCountAnd(x uint8) int {
	ans := 0
	for x != 0 {
		x = x & (x - 1)
		ans++
	}
	return ans
}
