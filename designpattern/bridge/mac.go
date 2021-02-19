package main

import "fmt"

type mac struct {
	print
}

func NewMac(p print) *mac {
	return &mac{print: p}
}

func (m *mac) ComputerPrint() {
	fmt.Print("mac: ")
	m.print.print()
}
