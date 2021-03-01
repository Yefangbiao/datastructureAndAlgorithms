package main

import "fmt"

type freightTrain struct {
	name string
	mediator
}

func (f *freightTrain) arrive() {
	fmt.Printf("%s is apply for depart!\n", f.name)
	if f.mediator.canArrive(f) {
		f.depart()
	} else {
		fmt.Println("Platform block, Waiting!")
	}
}

func (f *freightTrain) depart() {
	fmt.Printf("%s is leaving\n", f.name)
	f.mediator.notifyDepart()
}

func (f *freightTrain) permitArrival() {
	fmt.Printf("%s is permitted depart! Leaving now\n", f.name)
}
