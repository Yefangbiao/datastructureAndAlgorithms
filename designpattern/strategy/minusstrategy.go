package main

import "fmt"

type minusStrategy struct {
}

func (m *minusStrategy) execute(a, b int) {
	fmt.Printf("execute minus, result: %v\n", a-b)
}
