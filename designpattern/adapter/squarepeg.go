package main

type squarePeg struct {
	width float64
}

func (s *squarePeg) Width() float64 {
	return s.width
}
