package main

import "fmt"

type register struct {
	next department
}

func (r *register) setNext(next department) {
	r.next = next
}

func (r *register) execute(p *patient) {
	if p.isRegister {
		fmt.Printf("patient:%v has already register\n", p.name)
	} else {
		p.isRegister = true
		fmt.Printf("patient:%v is register now\n", p.name)
	}
	r.next.execute(p)
}
