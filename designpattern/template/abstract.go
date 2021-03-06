package main

import "fmt"

type iAbstract interface {
	step1()
	step2()
	step3()
	step4()
}

type abstract struct {
	iAbstract
}

func (a *abstract) Start() {
	fmt.Println("Welcome, let's start!")
}

func (a *abstract) End() {
	fmt.Println("This is End!")
}
