package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 从重复的生成器取出前10个
	done := make(chan interface{})
	defer close(done)

	//for num := range take(done, repeat(done, 1), 10) {
	//	fmt.Printf("%d ", num)
	//}

	for num := range take(done, repeatFn(done, mRand), 10) {
		fmt.Printf("%v ", num)
	}

}

func mRand() interface{} {
	return rand.Int()
}
