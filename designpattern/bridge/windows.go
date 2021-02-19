package main

import "fmt"

type windows struct {
	print
}

func (w *windows) ComputerPrint() {
	fmt.Print("windows: ")
	w.print.print()
}

func NewWindows(p print) *windows {
	return &windows{print: p}
}
