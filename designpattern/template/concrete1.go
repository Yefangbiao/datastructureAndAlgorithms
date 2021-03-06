package main

import "fmt"

type concrete1 struct {
}

func (c *concrete1) step1() {
	fmt.Println("This is concrete1's step1")
}

func (c *concrete1) step2() {
	fmt.Println("This is concrete1's step2")
}

func (c *concrete1) step3() {
	fmt.Println("This is concrete1's step3")
}

func (c *concrete1) step4() {
	fmt.Println("This is concrete1's step4")
}
