package main

import "fmt"

type checkout struct {
	next department
}

func (c *checkout) setNext(next department) {
	c.next = next
}

func (c *checkout) execute(p *patient) {
	if p.isCheckOut {
		fmt.Printf("patient:%v has already isCheckOut\n", p.name)
	} else {
		p.isCheckOut = true
		fmt.Printf("patient:%v is isCheckOut now\n", p.name)
	}
	c.next.execute(p)
}
