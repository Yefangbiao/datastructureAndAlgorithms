package main

import "math"

type squareAdapter struct {
	*squarePeg
}

func (s *squareAdapter) getRadius() float64 {
	return s.Width() * math.Sqrt(2) / 2
}
