package main

type potato struct {
	price int
}

func (p *potato) GetPrice() int {
	return p.price
}
