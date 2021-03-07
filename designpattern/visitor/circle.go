package main

type circle struct {
	x, y   int
	radius float64
}

func (c *circle) accept(v visitor) {
	v.visitCircle(c)
}
