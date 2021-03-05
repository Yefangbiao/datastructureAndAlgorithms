package main

import "fmt"

type addStrategy struct {
}

func (a2 *addStrategy) execute(a, b int) {
	fmt.Printf("execute add, result: %v\n", a+b)
}
