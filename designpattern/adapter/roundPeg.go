package main

type roundPeg struct {
	radius float64
}

func (r *roundPeg) getRadius() float64 {
	return r.radius
}
