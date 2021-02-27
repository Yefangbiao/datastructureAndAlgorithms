package main

import "fmt"

type tvDevice struct {
	isOpen bool
}

func (t *tvDevice) on() {
	t.isOpen = true
	fmt.Println("Tv is on")
}

func (t *tvDevice) off() {
	t.isOpen = false
	fmt.Println("Tv is off")
}
