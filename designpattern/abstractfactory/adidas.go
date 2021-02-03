package main

import "fmt"

type adidas struct {
}

func (a adidas) newShirt(size int) shirt {
	return adidasShirt{iShirt: iShirt{
		producer: "adidas",
		size:     size,
	}}
}

func (a adidas) newShoe(size int) shoe {
	return &adidasShoe{iShoe: iShoe{
		producer: "adidas",
		size:     size,
	}}
}

type adidasShoe struct {
	iShoe
}

func (a adidasShoe) Action() {
	fmt.Println("穿上了adidas鞋")
}

type adidasShirt struct {
	iShirt
}

func (a adidasShirt) Action() {
	fmt.Println("穿上了adidas衣服")
}
