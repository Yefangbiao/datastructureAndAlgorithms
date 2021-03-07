package main

import "fmt"

type concreteVisitor struct {
}

func (c concreteVisitor) visitCircle(c2 *circle) {
	fmt.Printf("circle: point:(%d, %d), raduis:%.3f\n", c2.x, c2.y, c2.radius)
}

func (c concreteVisitor) visitorSquare(s *square) {
	fmt.Printf("square: length:%d, width:%d\n", s.length, s.width)
}
