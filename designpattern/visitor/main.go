package main

func main() {
	c := circle{
		x:      3,
		y:      4,
		radius: 324.23,
	}

	s := square{
		length: 234,
		width:  26,
	}

	v := concreteVisitor{}

	c.accept(v)
	s.accept(v)
}
