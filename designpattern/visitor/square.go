package main

type square struct {
	length int
	width  int
}

func (s *square) accept(v visitor) {
	v.visitorSquare(s)
}
