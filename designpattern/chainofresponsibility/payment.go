package main

import "fmt"

type payment struct {
	next department
}

func (p *payment) setNext(next department) {
	p.next = next
}

func (p *payment) execute(patient *patient) {
	if patient.isPayment {
		fmt.Printf("patient:%v has already isPayment\n", patient.name)
	} else {
		patient.isPayment = true
		fmt.Printf("patient:%v is isPayment now\n", patient.name)
	}
}
