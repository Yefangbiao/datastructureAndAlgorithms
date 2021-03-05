package main

func main() {
	c := newContext()

	addS := &addStrategy{}
	c.SetS(addS)
	c.executeStrategy(1, 3)

	minusS := &minusStrategy{}
	c.SetS(minusS)
	c.executeStrategy(5, 4)
}
