package main

import "fmt"

func main() {
	director := NewDirector()

	// normalHouse
	normalH := NewNormalHouse()
	director.SetBuilder(normalH)
	h1 := director.MakeHouse()

	fmt.Println(h1.door)
	fmt.Println(h1.window)
	fmt.Println(h1.wall)

	//ideologyHouse
	ideologyH := NewIdeologyHose()
	director.SetBuilder(ideologyH)
	h2 := director.MakeHouse()

	fmt.Println(h2.door)
	fmt.Println(h2.window)
	fmt.Println(h2.wall)

}
