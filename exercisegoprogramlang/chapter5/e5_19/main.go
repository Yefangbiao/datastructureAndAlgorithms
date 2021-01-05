package main

import "fmt"

func main() {
	fmt.Println(f())
}

type bailout struct{}

func f() (ans int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			//没有宕机
			return
		case bailout{}:
			ans = 1
		default:
			panic(p)
		}
	}()
	panic(bailout{})
}
