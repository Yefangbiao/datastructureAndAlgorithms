package main

import "fmt"

type hpPrint struct {
}

func (h *hpPrint) print() {
	fmt.Println("print by hp")
}
