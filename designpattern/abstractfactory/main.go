package main

import "fmt"

func main() {
	nikeFactory, err := GetFactory("nike")
	if err != nil {
		panic(err)
	}
	nikeShirt := nikeFactory.newShirt(3)
	nikeShoe := nikeFactory.newShoe(33)
	PrintShirt(nikeShirt)
	PrintShoe(nikeShoe)

	adidasFactory, err := GetFactory("adidas")
	if err != nil {
		panic(err)
	}
	adidasShirt := adidasFactory.newShirt(23)
	adidasShoe := adidasFactory.newShirt(3245)
	PrintShirt(adidasShirt)
	PrintShoe(adidasShoe)
}

func PrintShirt(s shirt) {
	s.Action()
	fmt.Printf("producer:%v size:%v\n", s.GetProducer(), s.GetSize())
}

func PrintShoe(s shoe) {
	s.Action()
	fmt.Printf("producer:%v size:%v\n", s.GetProducer(), s.GetSize())
}
