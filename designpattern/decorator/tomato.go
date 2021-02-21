package main

type tomato struct {
	price int
}

func (t *tomato) GetPrice() int {
	return t.price
}
