package main

import "fmt"

type publishState struct {
	*document
}

func (p publishState) execute() {
	fmt.Println("now is publish state, finish!")
}
