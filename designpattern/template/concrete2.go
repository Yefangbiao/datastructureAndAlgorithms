package main

import "fmt"

type concrete2 struct {
}

func (c *concrete2) step1() {
	fmt.Println("This is concrete2's step1")
}

func (c *concrete2) step2() {
	fmt.Println("This is concrete2's step1")
}

func (c *concrete2) step3() {
	fmt.Println("This is concrete2's step1")
}

func (c *concrete2) step4() {
	fmt.Println("This is concrete2's step1")
}
