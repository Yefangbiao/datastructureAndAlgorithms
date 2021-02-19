package main

import "fmt"

type epsonPrint struct {
}

func (e *epsonPrint) print() {
	fmt.Print("print by epson")
}
