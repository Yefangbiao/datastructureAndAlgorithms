package main

import "fmt"

type topChief struct {
	m []material
}

func (t *topChief) GetPrice() int {
	materialPrice := 0
	for _, m := range t.m {
		materialPrice += m.GetPrice()
	}
	return materialPrice + 10
}

func (t *topChief) cook() {
	fmt.Printf("顶尖名厨炒了一道菜, 价格:%v\n", t.GetPrice())
}
