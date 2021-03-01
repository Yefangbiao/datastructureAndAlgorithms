package main

import "fmt"

type passengerTrain struct {
	name string
	mediator
}

func (p *passengerTrain) arrive() {
	fmt.Printf("%s is apply for depart!\n", p.name)
	if p.mediator.canArrive(p) {
		p.depart()
	} else {
		fmt.Println("Platform block, Waiting!")
	}
}

func (p *passengerTrain) depart() {
	fmt.Printf("%s is leaving\n", p.name)
	p.mediator.notifyDepart()
}

func (p *passengerTrain) permitArrival() {
	fmt.Printf("%s is permitted depart! Leaving now\n", p.name)
}
