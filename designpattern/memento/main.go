package main

import "fmt"

func main() {
	origin := originator{}
	ct := caretaker{[]*memento{}}

	origin.SetState("A")
	fmt.Println("State: ", origin.State())
	ct.addMemento(origin.createMemento())

	origin.SetState("B")
	fmt.Println("State: ", origin.State())
	ct.addMemento(origin.createMemento())

	origin.restoreFromMemento(ct.getMemento())
	fmt.Println("State: ", origin.State())

	origin.restoreFromMemento(ct.getMemento())
	fmt.Println("State: ", origin.State())
}
