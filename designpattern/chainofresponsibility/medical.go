package main

import "fmt"

type medical struct {
	next department
}

func (m *medical) setNext(next department) {
	m.next = next
}

func (m *medical) execute(p *patient) {
	if p.isMedical {
		fmt.Printf("patient:%v has already medical\n", p.name)
	} else {
		p.isMedical = true
		fmt.Printf("patient:%v is medical now\n", p.name)
	}
	m.next.execute(p)
}
